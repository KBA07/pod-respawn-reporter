package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNotifRestart) DeepCopyInto(out *PodNotifRestart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNotifRestart.
func (in *PodNotifRestart) DeepCopy() *PodNotifRestart {
	if in == nil {
		return nil
	}
	out := new(PodNotifRestart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodNotifRestart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNotifRestartList) DeepCopyInto(out *PodNotifRestartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodNotifRestart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNotifRestartList.
func (in *PodNotifRestartList) DeepCopy() *PodNotifRestartList {
	if in == nil {
		return nil
	}
	out := new(PodNotifRestartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodNotifRestartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNotifRestartSpec) DeepCopyInto(out *PodNotifRestartSpec) {
	*out = *in
	if in.NamespacesToMonitor != nil {
		in, out := &in.NamespacesToMonitor, &out.NamespacesToMonitor
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNotifRestartSpec.
func (in *PodNotifRestartSpec) DeepCopy() *PodNotifRestartSpec {
	if in == nil {
		return nil
	}
	out := new(PodNotifRestartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNotifRestartStatus) DeepCopyInto(out *PodNotifRestartStatus) {
	*out = *in
	in.LastChecked.DeepCopyInto(&out.LastChecked)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNotifRestartStatus.
func (in *PodNotifRestartStatus) DeepCopy() *PodNotifRestartStatus {
	if in == nil {
		return nil
	}
	out := new(PodNotifRestartStatus)
	in.DeepCopyInto(out)
	return out
}