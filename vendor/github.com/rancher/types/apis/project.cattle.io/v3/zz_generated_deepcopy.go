package v3

import (
	reflect "reflect"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DeployConfig).DeepCopyInto(out.(*DeployConfig))
			return nil
		}, InType: reflect.TypeOf(&DeployConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DeployStrategy).DeepCopyInto(out.(*DeployStrategy))
			return nil
		}, InType: reflect.TypeOf(&DeployStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DeploymentGlobalConfig).DeepCopyInto(out.(*DeploymentGlobalConfig))
			return nil
		}, InType: reflect.TypeOf(&DeploymentGlobalConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DeploymentJobConfig).DeepCopyInto(out.(*DeploymentJobConfig))
			return nil
		}, InType: reflect.TypeOf(&DeploymentJobConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DeploymentOrderedConfig).DeepCopyInto(out.(*DeploymentOrderedConfig))
			return nil
		}, InType: reflect.TypeOf(&DeploymentOrderedConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DeploymentParallelConfig).DeepCopyInto(out.(*DeploymentParallelConfig))
			return nil
		}, InType: reflect.TypeOf(&DeploymentParallelConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Link).DeepCopyInto(out.(*Link))
			return nil
		}, InType: reflect.TypeOf(&Link{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Workload).DeepCopyInto(out.(*Workload))
			return nil
		}, InType: reflect.TypeOf(&Workload{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkloadList).DeepCopyInto(out.(*WorkloadList))
			return nil
		}, InType: reflect.TypeOf(&WorkloadList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkloadSpec).DeepCopyInto(out.(*WorkloadSpec))
			return nil
		}, InType: reflect.TypeOf(&WorkloadSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkloadStatus).DeepCopyInto(out.(*WorkloadStatus))
			return nil
		}, InType: reflect.TypeOf(&WorkloadStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployConfig) DeepCopyInto(out *DeployConfig) {
	*out = *in
	if in.DeploymentStrategy != nil {
		in, out := &in.DeploymentStrategy, &out.DeploymentStrategy
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeployStrategy)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployConfig.
func (in *DeployConfig) DeepCopy() *DeployConfig {
	if in == nil {
		return nil
	}
	out := new(DeployConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployStrategy) DeepCopyInto(out *DeployStrategy) {
	*out = *in
	if in.ParallelConfig != nil {
		in, out := &in.ParallelConfig, &out.ParallelConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeploymentParallelConfig)
			**out = **in
		}
	}
	if in.JobConfig != nil {
		in, out := &in.JobConfig, &out.JobConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeploymentJobConfig)
			**out = **in
		}
	}
	if in.OrderedConfig != nil {
		in, out := &in.OrderedConfig, &out.OrderedConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeploymentOrderedConfig)
			**out = **in
		}
	}
	if in.GlobalConfig != nil {
		in, out := &in.GlobalConfig, &out.GlobalConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeploymentGlobalConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployStrategy.
func (in *DeployStrategy) DeepCopy() *DeployStrategy {
	if in == nil {
		return nil
	}
	out := new(DeployStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGlobalConfig) DeepCopyInto(out *DeploymentGlobalConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGlobalConfig.
func (in *DeploymentGlobalConfig) DeepCopy() *DeploymentGlobalConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentGlobalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentJobConfig) DeepCopyInto(out *DeploymentJobConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentJobConfig.
func (in *DeploymentJobConfig) DeepCopy() *DeploymentJobConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentJobConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentOrderedConfig) DeepCopyInto(out *DeploymentOrderedConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentOrderedConfig.
func (in *DeploymentOrderedConfig) DeepCopy() *DeploymentOrderedConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentOrderedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentParallelConfig) DeepCopyInto(out *DeploymentParallelConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentParallelConfig.
func (in *DeploymentParallelConfig) DeepCopy() *DeploymentParallelConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentParallelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Link) DeepCopyInto(out *Link) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(WorkloadStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadList) DeepCopyInto(out *WorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadList.
func (in *WorkloadList) DeepCopy() *WorkloadList {
	if in == nil {
		return nil
	}
	out := new(WorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSpec) DeepCopyInto(out *WorkloadSpec) {
	*out = *in
	in.DeployConfig.DeepCopyInto(&out.DeployConfig)
	in.Template.DeepCopyInto(&out.Template)
	if in.ServiceLinks != nil {
		in, out := &in.ServiceLinks, &out.ServiceLinks
		*out = make([]Link, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSpec.
func (in *WorkloadSpec) DeepCopy() *WorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}
