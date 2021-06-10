// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FalconAPI) DeepCopyInto(out *FalconAPI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FalconAPI.
func (in *FalconAPI) DeepCopy() *FalconAPI {
	if in == nil {
		return nil
	}
	out := new(FalconAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FalconConfig) DeepCopyInto(out *FalconConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FalconConfig.
func (in *FalconConfig) DeepCopy() *FalconConfig {
	if in == nil {
		return nil
	}
	out := new(FalconConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FalconConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FalconConfigList) DeepCopyInto(out *FalconConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FalconConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FalconConfigList.
func (in *FalconConfigList) DeepCopy() *FalconConfigList {
	if in == nil {
		return nil
	}
	out := new(FalconConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FalconConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FalconConfigSpec) DeepCopyInto(out *FalconConfigSpec) {
	*out = *in
	out.FalconAPI = in.FalconAPI
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FalconConfigSpec.
func (in *FalconConfigSpec) DeepCopy() *FalconConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FalconConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FalconConfigStatus) DeepCopyInto(out *FalconConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FalconConfigStatus.
func (in *FalconConfigStatus) DeepCopy() *FalconConfigStatus {
	if in == nil {
		return nil
	}
	out := new(FalconConfigStatus)
	in.DeepCopyInto(out)
	return out
}
