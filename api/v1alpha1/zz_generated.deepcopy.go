//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 SPIRE Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleEndpointProfile) DeepCopyInto(out *BundleEndpointProfile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleEndpointProfile.
func (in *BundleEndpointProfile) DeepCopy() *BundleEndpointProfile {
	if in == nil {
		return nil
	}
	out := new(BundleEndpointProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFederatedTrustDomain) DeepCopyInto(out *ClusterFederatedTrustDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFederatedTrustDomain.
func (in *ClusterFederatedTrustDomain) DeepCopy() *ClusterFederatedTrustDomain {
	if in == nil {
		return nil
	}
	out := new(ClusterFederatedTrustDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterFederatedTrustDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFederatedTrustDomainList) DeepCopyInto(out *ClusterFederatedTrustDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterFederatedTrustDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFederatedTrustDomainList.
func (in *ClusterFederatedTrustDomainList) DeepCopy() *ClusterFederatedTrustDomainList {
	if in == nil {
		return nil
	}
	out := new(ClusterFederatedTrustDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterFederatedTrustDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFederatedTrustDomainSpec) DeepCopyInto(out *ClusterFederatedTrustDomainSpec) {
	*out = *in
	out.BundleEndpointProfile = in.BundleEndpointProfile
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFederatedTrustDomainSpec.
func (in *ClusterFederatedTrustDomainSpec) DeepCopy() *ClusterFederatedTrustDomainSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterFederatedTrustDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFederatedTrustDomainStatus) DeepCopyInto(out *ClusterFederatedTrustDomainStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFederatedTrustDomainStatus.
func (in *ClusterFederatedTrustDomainStatus) DeepCopy() *ClusterFederatedTrustDomainStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterFederatedTrustDomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSPIFFEID) DeepCopyInto(out *ClusterSPIFFEID) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSPIFFEID.
func (in *ClusterSPIFFEID) DeepCopy() *ClusterSPIFFEID {
	if in == nil {
		return nil
	}
	out := new(ClusterSPIFFEID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSPIFFEID) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSPIFFEIDList) DeepCopyInto(out *ClusterSPIFFEIDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSPIFFEID, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSPIFFEIDList.
func (in *ClusterSPIFFEIDList) DeepCopy() *ClusterSPIFFEIDList {
	if in == nil {
		return nil
	}
	out := new(ClusterSPIFFEIDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSPIFFEIDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSPIFFEIDSpec) DeepCopyInto(out *ClusterSPIFFEIDSpec) {
	*out = *in
	out.TTL = in.TTL
	if in.DNSNameTemplates != nil {
		in, out := &in.DNSNameTemplates, &out.DNSNameTemplates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkloadSelectorTemplates != nil {
		in, out := &in.WorkloadSelectorTemplates, &out.WorkloadSelectorTemplates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FederatesWith != nil {
		in, out := &in.FederatesWith, &out.FederatesWith
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSPIFFEIDSpec.
func (in *ClusterSPIFFEIDSpec) DeepCopy() *ClusterSPIFFEIDSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSPIFFEIDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSPIFFEIDStats) DeepCopyInto(out *ClusterSPIFFEIDStats) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSPIFFEIDStats.
func (in *ClusterSPIFFEIDStats) DeepCopy() *ClusterSPIFFEIDStats {
	if in == nil {
		return nil
	}
	out := new(ClusterSPIFFEIDStats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSPIFFEIDStatus) DeepCopyInto(out *ClusterSPIFFEIDStatus) {
	*out = *in
	out.Stats = in.Stats
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSPIFFEIDStatus.
func (in *ClusterSPIFFEIDStatus) DeepCopy() *ClusterSPIFFEIDStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSPIFFEIDStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfig) DeepCopyInto(out *ProjectConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	if in.IgnoreNamespaces != nil {
		in, out := &in.IgnoreNamespaces, &out.IgnoreNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfig.
func (in *ProjectConfig) DeepCopy() *ProjectConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
