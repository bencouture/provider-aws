//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZone) DeepCopyInto(out *AvailabilityZone) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZone.
func (in *AvailabilityZone) DeepCopy() *AvailabilityZone {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Broker) DeepCopyInto(out *Broker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Broker.
func (in *Broker) DeepCopy() *Broker {
	if in == nil {
		return nil
	}
	out := new(Broker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Broker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerEngineType) DeepCopyInto(out *BrokerEngineType) {
	*out = *in
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerEngineType.
func (in *BrokerEngineType) DeepCopy() *BrokerEngineType {
	if in == nil {
		return nil
	}
	out := new(BrokerEngineType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerInstance) DeepCopyInto(out *BrokerInstance) {
	*out = *in
	if in.ConsoleURL != nil {
		in, out := &in.ConsoleURL, &out.ConsoleURL
		*out = new(string)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerInstance.
func (in *BrokerInstance) DeepCopy() *BrokerInstance {
	if in == nil {
		return nil
	}
	out := new(BrokerInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerInstanceOption) DeepCopyInto(out *BrokerInstanceOption) {
	*out = *in
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.HostInstanceType != nil {
		in, out := &in.HostInstanceType, &out.HostInstanceType
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SupportedEngineVersions != nil {
		in, out := &in.SupportedEngineVersions, &out.SupportedEngineVersions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerInstanceOption.
func (in *BrokerInstanceOption) DeepCopy() *BrokerInstanceOption {
	if in == nil {
		return nil
	}
	out := new(BrokerInstanceOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerList) DeepCopyInto(out *BrokerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Broker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerList.
func (in *BrokerList) DeepCopy() *BrokerList {
	if in == nil {
		return nil
	}
	out := new(BrokerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BrokerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerObservation) DeepCopyInto(out *BrokerObservation) {
	*out = *in
	if in.BrokerARN != nil {
		in, out := &in.BrokerARN, &out.BrokerARN
		*out = new(string)
		**out = **in
	}
	if in.BrokerID != nil {
		in, out := &in.BrokerID, &out.BrokerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerObservation.
func (in *BrokerObservation) DeepCopy() *BrokerObservation {
	if in == nil {
		return nil
	}
	out := new(BrokerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerParameters) DeepCopyInto(out *BrokerParameters) {
	*out = *in
	if in.AuthenticationStrategy != nil {
		in, out := &in.AuthenticationStrategy, &out.AuthenticationStrategy
		*out = new(string)
		**out = **in
	}
	if in.AutoMinorVersionUpgrade != nil {
		in, out := &in.AutoMinorVersionUpgrade, &out.AutoMinorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(ConfigurationID)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatorRequestID != nil {
		in, out := &in.CreatorRequestID, &out.CreatorRequestID
		*out = new(string)
		**out = **in
	}
	if in.DeploymentMode != nil {
		in, out := &in.DeploymentMode, &out.DeploymentMode
		*out = new(string)
		**out = **in
	}
	if in.EncryptionOptions != nil {
		in, out := &in.EncryptionOptions, &out.EncryptionOptions
		*out = new(EncryptionOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.HostInstanceType != nil {
		in, out := &in.HostInstanceType, &out.HostInstanceType
		*out = new(string)
		**out = **in
	}
	if in.LDAPServerMetadata != nil {
		in, out := &in.LDAPServerMetadata, &out.LDAPServerMetadata
		*out = new(LDAPServerMetadataInput)
		(*in).DeepCopyInto(*out)
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = new(Logs)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceWindowStartTime != nil {
		in, out := &in.MaintenanceWindowStartTime, &out.MaintenanceWindowStartTime
		*out = new(WeeklyStartTime)
		(*in).DeepCopyInto(*out)
	}
	if in.PubliclyAccessible != nil {
		in, out := &in.PubliclyAccessible, &out.PubliclyAccessible
		*out = new(bool)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	in.CustomBrokerParameters.DeepCopyInto(&out.CustomBrokerParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerParameters.
func (in *BrokerParameters) DeepCopy() *BrokerParameters {
	if in == nil {
		return nil
	}
	out := new(BrokerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpec) DeepCopyInto(out *BrokerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpec.
func (in *BrokerSpec) DeepCopy() *BrokerSpec {
	if in == nil {
		return nil
	}
	out := new(BrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerStatus) DeepCopyInto(out *BrokerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerStatus.
func (in *BrokerStatus) DeepCopy() *BrokerStatus {
	if in == nil {
		return nil
	}
	out := new(BrokerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSummary) DeepCopyInto(out *BrokerSummary) {
	*out = *in
	if in.BrokerARN != nil {
		in, out := &in.BrokerARN, &out.BrokerARN
		*out = new(string)
		**out = **in
	}
	if in.BrokerID != nil {
		in, out := &in.BrokerID, &out.BrokerID
		*out = new(string)
		**out = **in
	}
	if in.BrokerName != nil {
		in, out := &in.BrokerName, &out.BrokerName
		*out = new(string)
		**out = **in
	}
	if in.BrokerState != nil {
		in, out := &in.BrokerState, &out.BrokerState
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.DeploymentMode != nil {
		in, out := &in.DeploymentMode, &out.DeploymentMode
		*out = new(string)
		**out = **in
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.HostInstanceType != nil {
		in, out := &in.HostInstanceType, &out.HostInstanceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSummary.
func (in *BrokerSummary) DeepCopy() *BrokerSummary {
	if in == nil {
		return nil
	}
	out := new(BrokerSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationStrategy != nil {
		in, out := &in.AuthenticationStrategy, &out.AuthenticationStrategy
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EngineType != nil {
		in, out := &in.EngineType, &out.EngineType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
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
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationID) DeepCopyInto(out *ConfigurationID) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationID.
func (in *ConfigurationID) DeepCopy() *ConfigurationID {
	if in == nil {
		return nil
	}
	out := new(ConfigurationID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationRevision) DeepCopyInto(out *ConfigurationRevision) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationRevision.
func (in *ConfigurationRevision) DeepCopy() *ConfigurationRevision {
	if in == nil {
		return nil
	}
	out := new(ConfigurationRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configurations) DeepCopyInto(out *Configurations) {
	*out = *in
	if in.Current != nil {
		in, out := &in.Current, &out.Current
		*out = new(ConfigurationID)
		(*in).DeepCopyInto(*out)
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]*ConfigurationID, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ConfigurationID)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = new(ConfigurationID)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configurations.
func (in *Configurations) DeepCopy() *Configurations {
	if in == nil {
		return nil
	}
	out := new(Configurations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomBrokerParameters) DeepCopyInto(out *CustomBrokerParameters) {
	*out = *in
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomUsers != nil {
		in, out := &in.CustomUsers, &out.CustomUsers
		*out = make([]*CustomUser, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CustomUser)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomBrokerParameters.
func (in *CustomBrokerParameters) DeepCopy() *CustomBrokerParameters {
	if in == nil {
		return nil
	}
	out := new(CustomBrokerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUser) DeepCopyInto(out *CustomUser) {
	*out = *in
	if in.ConsoleAccess != nil {
		in, out := &in.ConsoleAccess, &out.ConsoleAccess
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUser.
func (in *CustomUser) DeepCopy() *CustomUser {
	if in == nil {
		return nil
	}
	out := new(CustomUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUserParameters) DeepCopyInto(out *CustomUserParameters) {
	*out = *in
	if in.BrokerID != nil {
		in, out := &in.BrokerID, &out.BrokerID
		*out = new(string)
		**out = **in
	}
	if in.BrokerIDRef != nil {
		in, out := &in.BrokerIDRef, &out.BrokerIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.BrokerIDSelector != nil {
		in, out := &in.BrokerIDSelector, &out.BrokerIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	out.PasswordSecretRef = in.PasswordSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUserParameters.
func (in *CustomUserParameters) DeepCopy() *CustomUserParameters {
	if in == nil {
		return nil
	}
	out := new(CustomUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionOptions) DeepCopyInto(out *EncryptionOptions) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.UseAWSOwnedKey != nil {
		in, out := &in.UseAWSOwnedKey, &out.UseAWSOwnedKey
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionOptions.
func (in *EncryptionOptions) DeepCopy() *EncryptionOptions {
	if in == nil {
		return nil
	}
	out := new(EncryptionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EngineVersion) DeepCopyInto(out *EngineVersion) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EngineVersion.
func (in *EngineVersion) DeepCopy() *EngineVersion {
	if in == nil {
		return nil
	}
	out := new(EngineVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAPServerMetadataInput) DeepCopyInto(out *LDAPServerMetadataInput) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RoleBase != nil {
		in, out := &in.RoleBase, &out.RoleBase
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchMatching != nil {
		in, out := &in.RoleSearchMatching, &out.RoleSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchSubtree != nil {
		in, out := &in.RoleSearchSubtree, &out.RoleSearchSubtree
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccountPassword != nil {
		in, out := &in.ServiceAccountPassword, &out.ServiceAccountPassword
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountUsername != nil {
		in, out := &in.ServiceAccountUsername, &out.ServiceAccountUsername
		*out = new(string)
		**out = **in
	}
	if in.UserBase != nil {
		in, out := &in.UserBase, &out.UserBase
		*out = new(string)
		**out = **in
	}
	if in.UserRoleName != nil {
		in, out := &in.UserRoleName, &out.UserRoleName
		*out = new(string)
		**out = **in
	}
	if in.UserSearchMatching != nil {
		in, out := &in.UserSearchMatching, &out.UserSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.UserSearchSubtree != nil {
		in, out := &in.UserSearchSubtree, &out.UserSearchSubtree
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAPServerMetadataInput.
func (in *LDAPServerMetadataInput) DeepCopy() *LDAPServerMetadataInput {
	if in == nil {
		return nil
	}
	out := new(LDAPServerMetadataInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAPServerMetadataOutput) DeepCopyInto(out *LDAPServerMetadataOutput) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RoleBase != nil {
		in, out := &in.RoleBase, &out.RoleBase
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchMatching != nil {
		in, out := &in.RoleSearchMatching, &out.RoleSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.RoleSearchSubtree != nil {
		in, out := &in.RoleSearchSubtree, &out.RoleSearchSubtree
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccountUsername != nil {
		in, out := &in.ServiceAccountUsername, &out.ServiceAccountUsername
		*out = new(string)
		**out = **in
	}
	if in.UserBase != nil {
		in, out := &in.UserBase, &out.UserBase
		*out = new(string)
		**out = **in
	}
	if in.UserRoleName != nil {
		in, out := &in.UserRoleName, &out.UserRoleName
		*out = new(string)
		**out = **in
	}
	if in.UserSearchMatching != nil {
		in, out := &in.UserSearchMatching, &out.UserSearchMatching
		*out = new(string)
		**out = **in
	}
	if in.UserSearchSubtree != nil {
		in, out := &in.UserSearchSubtree, &out.UserSearchSubtree
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAPServerMetadataOutput.
func (in *LDAPServerMetadataOutput) DeepCopy() *LDAPServerMetadataOutput {
	if in == nil {
		return nil
	}
	out := new(LDAPServerMetadataOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logs) DeepCopyInto(out *Logs) {
	*out = *in
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(bool)
		**out = **in
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logs.
func (in *Logs) DeepCopy() *Logs {
	if in == nil {
		return nil
	}
	out := new(Logs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsSummary) DeepCopyInto(out *LogsSummary) {
	*out = *in
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(bool)
		**out = **in
	}
	if in.AuditLogGroup != nil {
		in, out := &in.AuditLogGroup, &out.AuditLogGroup
		*out = new(string)
		**out = **in
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(bool)
		**out = **in
	}
	if in.GeneralLogGroup != nil {
		in, out := &in.GeneralLogGroup, &out.GeneralLogGroup
		*out = new(string)
		**out = **in
	}
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = new(PendingLogs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsSummary.
func (in *LogsSummary) DeepCopy() *LogsSummary {
	if in == nil {
		return nil
	}
	out := new(LogsSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingLogs) DeepCopyInto(out *PendingLogs) {
	*out = *in
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(bool)
		**out = **in
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingLogs.
func (in *PendingLogs) DeepCopy() *PendingLogs {
	if in == nil {
		return nil
	}
	out := new(PendingLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SanitizationWarning) DeepCopyInto(out *SanitizationWarning) {
	*out = *in
	if in.AttributeName != nil {
		in, out := &in.AttributeName, &out.AttributeName
		*out = new(string)
		**out = **in
	}
	if in.ElementName != nil {
		in, out := &in.ElementName, &out.ElementName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SanitizationWarning.
func (in *SanitizationWarning) DeepCopy() *SanitizationWarning {
	if in == nil {
		return nil
	}
	out := new(SanitizationWarning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserObservation) DeepCopyInto(out *UserObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserObservation.
func (in *UserObservation) DeepCopy() *UserObservation {
	if in == nil {
		return nil
	}
	out := new(UserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserParameters) DeepCopyInto(out *UserParameters) {
	*out = *in
	if in.ConsoleAccess != nil {
		in, out := &in.ConsoleAccess, &out.ConsoleAccess
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	in.CustomUserParameters.DeepCopyInto(&out.CustomUserParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserParameters.
func (in *UserParameters) DeepCopy() *UserParameters {
	if in == nil {
		return nil
	}
	out := new(UserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPendingChanges) DeepCopyInto(out *UserPendingChanges) {
	*out = *in
	if in.ConsoleAccess != nil {
		in, out := &in.ConsoleAccess, &out.ConsoleAccess
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PendingChange != nil {
		in, out := &in.PendingChange, &out.PendingChange
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPendingChanges.
func (in *UserPendingChanges) DeepCopy() *UserPendingChanges {
	if in == nil {
		return nil
	}
	out := new(UserPendingChanges)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSummary) DeepCopyInto(out *UserSummary) {
	*out = *in
	if in.PendingChange != nil {
		in, out := &in.PendingChange, &out.PendingChange
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSummary.
func (in *UserSummary) DeepCopy() *UserSummary {
	if in == nil {
		return nil
	}
	out := new(UserSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User_SDK) DeepCopyInto(out *User_SDK) {
	*out = *in
	if in.ConsoleAccess != nil {
		in, out := &in.ConsoleAccess, &out.ConsoleAccess
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User_SDK.
func (in *User_SDK) DeepCopy() *User_SDK {
	if in == nil {
		return nil
	}
	out := new(User_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStartTime) DeepCopyInto(out *WeeklyStartTime) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.TimeOfDay != nil {
		in, out := &in.TimeOfDay, &out.TimeOfDay
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStartTime.
func (in *WeeklyStartTime) DeepCopy() *WeeklyStartTime {
	if in == nil {
		return nil
	}
	out := new(WeeklyStartTime)
	in.DeepCopyInto(out)
	return out
}
