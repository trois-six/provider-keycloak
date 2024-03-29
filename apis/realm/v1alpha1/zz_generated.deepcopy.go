//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Events) DeepCopyInto(out *Events) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Events.
func (in *Events) DeepCopy() *Events {
	if in == nil {
		return nil
	}
	out := new(Events)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Events) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsInitParameters) DeepCopyInto(out *EventsInitParameters) {
	*out = *in
	if in.AdminEventsDetailsEnabled != nil {
		in, out := &in.AdminEventsDetailsEnabled, &out.AdminEventsDetailsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AdminEventsEnabled != nil {
		in, out := &in.AdminEventsEnabled, &out.AdminEventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EnabledEventTypes != nil {
		in, out := &in.EnabledEventTypes, &out.EnabledEventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EventsEnabled != nil {
		in, out := &in.EventsEnabled, &out.EventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EventsExpiration != nil {
		in, out := &in.EventsExpiration, &out.EventsExpiration
		*out = new(float64)
		**out = **in
	}
	if in.EventsListeners != nil {
		in, out := &in.EventsListeners, &out.EventsListeners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsInitParameters.
func (in *EventsInitParameters) DeepCopy() *EventsInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsList) DeepCopyInto(out *EventsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Events, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsList.
func (in *EventsList) DeepCopy() *EventsList {
	if in == nil {
		return nil
	}
	out := new(EventsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsObservation) DeepCopyInto(out *EventsObservation) {
	*out = *in
	if in.AdminEventsDetailsEnabled != nil {
		in, out := &in.AdminEventsDetailsEnabled, &out.AdminEventsDetailsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AdminEventsEnabled != nil {
		in, out := &in.AdminEventsEnabled, &out.AdminEventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EnabledEventTypes != nil {
		in, out := &in.EnabledEventTypes, &out.EnabledEventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EventsEnabled != nil {
		in, out := &in.EventsEnabled, &out.EventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EventsExpiration != nil {
		in, out := &in.EventsExpiration, &out.EventsExpiration
		*out = new(float64)
		**out = **in
	}
	if in.EventsListeners != nil {
		in, out := &in.EventsListeners, &out.EventsListeners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsObservation.
func (in *EventsObservation) DeepCopy() *EventsObservation {
	if in == nil {
		return nil
	}
	out := new(EventsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsParameters) DeepCopyInto(out *EventsParameters) {
	*out = *in
	if in.AdminEventsDetailsEnabled != nil {
		in, out := &in.AdminEventsDetailsEnabled, &out.AdminEventsDetailsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AdminEventsEnabled != nil {
		in, out := &in.AdminEventsEnabled, &out.AdminEventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EnabledEventTypes != nil {
		in, out := &in.EnabledEventTypes, &out.EnabledEventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EventsEnabled != nil {
		in, out := &in.EventsEnabled, &out.EventsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EventsExpiration != nil {
		in, out := &in.EventsExpiration, &out.EventsExpiration
		*out = new(float64)
		**out = **in
	}
	if in.EventsListeners != nil {
		in, out := &in.EventsListeners, &out.EventsListeners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsParameters.
func (in *EventsParameters) DeepCopy() *EventsParameters {
	if in == nil {
		return nil
	}
	out := new(EventsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsSpec) DeepCopyInto(out *EventsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsSpec.
func (in *EventsSpec) DeepCopy() *EventsSpec {
	if in == nil {
		return nil
	}
	out := new(EventsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsStatus) DeepCopyInto(out *EventsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsStatus.
func (in *EventsStatus) DeepCopy() *EventsStatus {
	if in == nil {
		return nil
	}
	out := new(EventsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmKeyStoreAESGenerated) DeepCopyInto(out *RealmKeyStoreAESGenerated) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmKeyStoreAESGenerated.
func (in *RealmKeyStoreAESGenerated) DeepCopy() *RealmKeyStoreAESGenerated {
	if in == nil {
		return nil
	}
	out := new(RealmKeyStoreAESGenerated)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealmKeyStoreAESGenerated) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmKeyStoreAESGeneratedInitParameters) DeepCopyInto(out *RealmKeyStoreAESGeneratedInitParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.SecretSize != nil {
		in, out := &in.SecretSize, &out.SecretSize
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmKeyStoreAESGeneratedInitParameters.
func (in *RealmKeyStoreAESGeneratedInitParameters) DeepCopy() *RealmKeyStoreAESGeneratedInitParameters {
	if in == nil {
		return nil
	}
	out := new(RealmKeyStoreAESGeneratedInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmKeyStoreAESGeneratedList) DeepCopyInto(out *RealmKeyStoreAESGeneratedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RealmKeyStoreAESGenerated, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmKeyStoreAESGeneratedList.
func (in *RealmKeyStoreAESGeneratedList) DeepCopy() *RealmKeyStoreAESGeneratedList {
	if in == nil {
		return nil
	}
	out := new(RealmKeyStoreAESGeneratedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealmKeyStoreAESGeneratedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmKeyStoreAESGeneratedObservation) DeepCopyInto(out *RealmKeyStoreAESGeneratedObservation) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.SecretSize != nil {
		in, out := &in.SecretSize, &out.SecretSize
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmKeyStoreAESGeneratedObservation.
func (in *RealmKeyStoreAESGeneratedObservation) DeepCopy() *RealmKeyStoreAESGeneratedObservation {
	if in == nil {
		return nil
	}
	out := new(RealmKeyStoreAESGeneratedObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmKeyStoreAESGeneratedParameters) DeepCopyInto(out *RealmKeyStoreAESGeneratedParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.SecretSize != nil {
		in, out := &in.SecretSize, &out.SecretSize
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmKeyStoreAESGeneratedParameters.
func (in *RealmKeyStoreAESGeneratedParameters) DeepCopy() *RealmKeyStoreAESGeneratedParameters {
	if in == nil {
		return nil
	}
	out := new(RealmKeyStoreAESGeneratedParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmKeyStoreAESGeneratedSpec) DeepCopyInto(out *RealmKeyStoreAESGeneratedSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmKeyStoreAESGeneratedSpec.
func (in *RealmKeyStoreAESGeneratedSpec) DeepCopy() *RealmKeyStoreAESGeneratedSpec {
	if in == nil {
		return nil
	}
	out := new(RealmKeyStoreAESGeneratedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealmKeyStoreAESGeneratedStatus) DeepCopyInto(out *RealmKeyStoreAESGeneratedStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealmKeyStoreAESGeneratedStatus.
func (in *RealmKeyStoreAESGeneratedStatus) DeepCopy() *RealmKeyStoreAESGeneratedStatus {
	if in == nil {
		return nil
	}
	out := new(RealmKeyStoreAESGeneratedStatus)
	in.DeepCopyInto(out)
	return out
}
