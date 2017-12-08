package authz

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/config"
	"github.com/sirupsen/logrus"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

const (
	CreateAction     = "create"
	UpdateAction     = "update"
	RemoveAction     = "remove"
	NoopAction       = "noop"
	name             = "binding-handler"
	bindingNameLabel = "io.cattle.field.projectRoleTemplateBindingName"
	projectIDLabel   = "io.cattle.field.projectId"
)

func Register(workload *config.ClusterContext) {
	r := &roleHandler{
		Namespaces:                 workload.K8sClient.CoreV1().Namespaces(),
		PodSecurityPolicies:        workload.K8sClient.ExtensionsV1beta1().PodSecurityPolicies(),
		RoleTemplates:              workload.Management.Management.RoleTemplates(""),
		PodSecurityPolicyTemplates: workload.Management.Management.PodSecurityPolicyTemplates(""),
		RBAC: workload.K8sClient.RbacV1(),
	}
	workload.Management.Management.ProjectRoleTemplateBindings("").Controller().AddHandler(r.sync)
}

type roleHandler struct {
	Namespaces                 v1.NamespaceInterface
	PodSecurityPolicies        v1beta1.PodSecurityPolicyInterface
	RoleTemplates              v3.RoleTemplateInterface
	PodSecurityPolicyTemplates v3.PodSecurityPolicyTemplateInterface
	RBAC                       rbacv1client.RbacV1Interface
}

func getAction(binding *v3.ProjectRoleTemplateBinding) string {
	if binding == nil {
		return NoopAction
	}

	if binding.ObjectMeta.DeletionTimestamp != nil {
		return RemoveAction
	}

	return CreateAction
}

func (r *roleHandler) sync(key string, binding *v3.ProjectRoleTemplateBinding) error {
	action := getAction(binding)
	switch action {
	case CreateAction:
		return r.create(key, binding)
	case UpdateAction:
		// TODO Handle update
		return nil
	case RemoveAction:
		// TODO Remove
	}

	return nil
}

func (r *roleHandler) create(key string, binding *v3.ProjectRoleTemplateBinding) error {
	rt, err := r.RoleTemplates.Get(binding.RoleTemplateName, metav1.GetOptions{})
	if err != nil {
		return errors.Wrapf(err, "couldn't get role template %v", binding.RoleTemplateName)
	}

	// Get namespaces belonging to project
	set := labels.Set(map[string]string{projectIDLabel: binding.ProjectName})
	namespaces, err := r.Namespaces.List(metav1.ListOptions{
		LabelSelector: set.AsSelector().String(),
	})
	if err != nil {
		return errors.Wrapf(err, "couldn't list namespaces with selector %s", set.AsSelector())
	}
	if len(namespaces.Items) == 0 {
		return nil
	}

	pspTemplates := map[string]*v3.PodSecurityPolicyTemplate{}
	roles := map[string]*v3.RoleTemplate{}

	if err := r.gatherRolesAndPSPs(rt, roles, pspTemplates); err != nil {
		return errors.Wrap(err, "couldn't gather RoleTemplates and PodSecurityPolicyTemplates")

	}
	if err := r.ensurePSPs(pspTemplates); err != nil {
		return errors.Wrap(err, "couldn't ensure PodSecurityPolicies")
	}

	if err := r.ensureRoles(roles); err != nil {
		return errors.Wrap(err, "couldn't ensure roles")
	}

	// TODO is .Items the complete list or is there potential pagination to deal with?
	for _, ns := range namespaces.Items {
		if err := r.ensureBinding(ns.Name, rt.Name, binding); err != nil {
			return errors.Wrapf(err, "couldn't ensure binding %v %v in %v", rt.Name, binding.Subject.Name, ns.Name)
		}
	}

	return nil
}

func (r *roleHandler) gatherRolesAndPSPs(rt *v3.RoleTemplate, roleTemplates map[string]*v3.RoleTemplate, pspTemplates map[string]*v3.PodSecurityPolicyTemplate) error {
	for _, pspName := range rt.PodSecurityPolicyTemplateNames {
		pspTemplate, err := r.PodSecurityPolicyTemplates.Get(pspName, metav1.GetOptions{})
		if err != nil {
			logrus.Warnf("Couldn't find PodSecurityPolicyTemplate %v. Skipping. Error: %v", pspName, err)
			continue
		}
		pspTemplates[pspName] = pspTemplate
	}

	roleTemplates[rt.Name] = rt

	for _, rtName := range rt.RoleTemplateNames {
		subRT, err := r.RoleTemplates.Get(rtName, metav1.GetOptions{})
		if err != nil {
			return errors.Wrapf(err, "couldn't get RoleTemplate %s", rtName)
		}
		if err := r.gatherRolesAndPSPs(subRT, roleTemplates, pspTemplates); err != nil {
			return errors.Wrapf(err, "couldn't gather RoleTemplate %s", rtName)
		}
	}

	return nil
}

func (r *roleHandler) ensurePSPs(pspTemplates map[string]*v3.PodSecurityPolicyTemplate) error {
	pspCli := r.PodSecurityPolicies
	for name, pspTemplate := range pspTemplates {
		if psp, err := pspCli.Get(name, metav1.GetOptions{}); err == nil {
			psp.Spec = pspTemplate.Spec
			if _, err := pspCli.Update(psp); err != nil {
				return errors.Wrapf(err, "couldn't update PodSecurityPolicy %v", name)
			}
			continue
		}
		_, err := pspCli.Create(&extv1beta1.PodSecurityPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: pspTemplate.Spec,
		})
		if err != nil {
			return errors.Wrapf(err, "couldn't create PodSecurityPolicy %v", name)
		}
	}

	return nil
}

func (r *roleHandler) ensureRoles(rts map[string]*v3.RoleTemplate) error {
	roleCli := r.RBAC.ClusterRoles()
	for _, rt := range rts {
		if rt.Builtin {
			// TODO assert the role exists and log an error if it doesnt.
			continue
		}

		if role, err := roleCli.Get(rt.Name, metav1.GetOptions{}); err == nil {
			// TODO potentially check a version so that we don't do unnecessary updates
			role.Rules = rt.Rules
			_, err := roleCli.Update(role)
			if err != nil {
				return errors.Wrapf(err, "couldn't update role %v", rt.Name)
			}
		}

		_, err := roleCli.Create(&rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name: rt.Name,
			},
			Rules: rt.Rules,
		})
		if err != nil {
			return errors.Wrapf(err, "couldn't create role %v", rt.Name)
		}
	}

	return nil
}

func (r *roleHandler) ensureBinding(ns, roleName string, binding *v3.ProjectRoleTemplateBinding) error {
	bindingCli := r.RBAC.RoleBindings(ns)
	bindingName := strings.ToLower(fmt.Sprintf("%v-%v-%v", roleName, binding.Subject.Kind, binding.Subject.Name))
	_, err := bindingCli.Get(bindingName, metav1.GetOptions{})
	if err == nil {
		return nil
	}

	_, err = bindingCli.Create(&rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:   bindingName,
			Labels: map[string]string{bindingNameLabel: binding.Name},
		},
		Subjects: []rbacv1.Subject{binding.Subject},
		RoleRef: rbacv1.RoleRef{
			Kind: "ClusterRole",
			Name: roleName,
		},
	})

	return err
}
