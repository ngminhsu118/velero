// +build !ignore_autogenerated

/*
Copyright 2017 the Heptio Ark contributors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AWSConfig).DeepCopyInto(out.(*AWSConfig))
			return nil
		}, InType: reflect.TypeOf(&AWSConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureConfig).DeepCopyInto(out.(*AzureConfig))
			return nil
		}, InType: reflect.TypeOf(&AzureConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Backup).DeepCopyInto(out.(*Backup))
			return nil
		}, InType: reflect.TypeOf(&Backup{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BackupHooks).DeepCopyInto(out.(*BackupHooks))
			return nil
		}, InType: reflect.TypeOf(&BackupHooks{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BackupList).DeepCopyInto(out.(*BackupList))
			return nil
		}, InType: reflect.TypeOf(&BackupList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BackupResourceHook).DeepCopyInto(out.(*BackupResourceHook))
			return nil
		}, InType: reflect.TypeOf(&BackupResourceHook{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BackupResourceHookSpec).DeepCopyInto(out.(*BackupResourceHookSpec))
			return nil
		}, InType: reflect.TypeOf(&BackupResourceHookSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BackupSpec).DeepCopyInto(out.(*BackupSpec))
			return nil
		}, InType: reflect.TypeOf(&BackupSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BackupStatus).DeepCopyInto(out.(*BackupStatus))
			return nil
		}, InType: reflect.TypeOf(&BackupStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CloudProviderConfig).DeepCopyInto(out.(*CloudProviderConfig))
			return nil
		}, InType: reflect.TypeOf(&CloudProviderConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Config).DeepCopyInto(out.(*Config))
			return nil
		}, InType: reflect.TypeOf(&Config{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ConfigList).DeepCopyInto(out.(*ConfigList))
			return nil
		}, InType: reflect.TypeOf(&ConfigList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DownloadRequest).DeepCopyInto(out.(*DownloadRequest))
			return nil
		}, InType: reflect.TypeOf(&DownloadRequest{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DownloadRequestList).DeepCopyInto(out.(*DownloadRequestList))
			return nil
		}, InType: reflect.TypeOf(&DownloadRequestList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DownloadRequestSpec).DeepCopyInto(out.(*DownloadRequestSpec))
			return nil
		}, InType: reflect.TypeOf(&DownloadRequestSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DownloadRequestStatus).DeepCopyInto(out.(*DownloadRequestStatus))
			return nil
		}, InType: reflect.TypeOf(&DownloadRequestStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DownloadTarget).DeepCopyInto(out.(*DownloadTarget))
			return nil
		}, InType: reflect.TypeOf(&DownloadTarget{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExecHook).DeepCopyInto(out.(*ExecHook))
			return nil
		}, InType: reflect.TypeOf(&ExecHook{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GCPConfig).DeepCopyInto(out.(*GCPConfig))
			return nil
		}, InType: reflect.TypeOf(&GCPConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ObjectStorageProviderConfig).DeepCopyInto(out.(*ObjectStorageProviderConfig))
			return nil
		}, InType: reflect.TypeOf(&ObjectStorageProviderConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Restore).DeepCopyInto(out.(*Restore))
			return nil
		}, InType: reflect.TypeOf(&Restore{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RestoreList).DeepCopyInto(out.(*RestoreList))
			return nil
		}, InType: reflect.TypeOf(&RestoreList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RestoreResult).DeepCopyInto(out.(*RestoreResult))
			return nil
		}, InType: reflect.TypeOf(&RestoreResult{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RestoreSpec).DeepCopyInto(out.(*RestoreSpec))
			return nil
		}, InType: reflect.TypeOf(&RestoreSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RestoreStatus).DeepCopyInto(out.(*RestoreStatus))
			return nil
		}, InType: reflect.TypeOf(&RestoreStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Schedule).DeepCopyInto(out.(*Schedule))
			return nil
		}, InType: reflect.TypeOf(&Schedule{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ScheduleList).DeepCopyInto(out.(*ScheduleList))
			return nil
		}, InType: reflect.TypeOf(&ScheduleList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ScheduleSpec).DeepCopyInto(out.(*ScheduleSpec))
			return nil
		}, InType: reflect.TypeOf(&ScheduleSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ScheduleStatus).DeepCopyInto(out.(*ScheduleStatus))
			return nil
		}, InType: reflect.TypeOf(&ScheduleStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeBackupInfo).DeepCopyInto(out.(*VolumeBackupInfo))
			return nil
		}, InType: reflect.TypeOf(&VolumeBackupInfo{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfig) DeepCopyInto(out *AWSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfig.
func (in *AWSConfig) DeepCopy() *AWSConfig {
	if in == nil {
		return nil
	}
	out := new(AWSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfig) DeepCopyInto(out *AzureConfig) {
	*out = *in
	out.APITimeout = in.APITimeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfig.
func (in *AzureConfig) DeepCopy() *AzureConfig {
	if in == nil {
		return nil
	}
	out := new(AzureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupHooks) DeepCopyInto(out *BackupHooks) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]BackupResourceHookSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupHooks.
func (in *BackupHooks) DeepCopy() *BackupHooks {
	if in == nil {
		return nil
	}
	out := new(BackupHooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupResourceHook) DeepCopyInto(out *BackupResourceHook) {
	*out = *in
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		if *in == nil {
			*out = nil
		} else {
			*out = new(ExecHook)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupResourceHook.
func (in *BackupResourceHook) DeepCopy() *BackupResourceHook {
	if in == nil {
		return nil
	}
	out := new(BackupResourceHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupResourceHookSpec) DeepCopyInto(out *BackupResourceHookSpec) {
	*out = *in
	if in.IncludedNamespaces != nil {
		in, out := &in.IncludedNamespaces, &out.IncludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedNamespaces != nil {
		in, out := &in.ExcludedNamespaces, &out.ExcludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludedResources != nil {
		in, out := &in.IncludedResources, &out.IncludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = make([]BackupResourceHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupResourceHookSpec.
func (in *BackupResourceHookSpec) DeepCopy() *BackupResourceHookSpec {
	if in == nil {
		return nil
	}
	out := new(BackupResourceHookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	if in.IncludedNamespaces != nil {
		in, out := &in.IncludedNamespaces, &out.IncludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedNamespaces != nil {
		in, out := &in.ExcludedNamespaces, &out.ExcludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludedResources != nil {
		in, out := &in.IncludedResources, &out.IncludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SnapshotVolumes != nil {
		in, out := &in.SnapshotVolumes, &out.SnapshotVolumes
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	out.TTL = in.TTL
	if in.IncludeClusterResources != nil {
		in, out := &in.IncludeClusterResources, &out.IncludeClusterResources
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	in.Hooks.DeepCopyInto(&out.Hooks)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	in.Expiration.DeepCopyInto(&out.Expiration)
	if in.VolumeBackups != nil {
		in, out := &in.VolumeBackups, &out.VolumeBackups
		*out = make(map[string]*VolumeBackupInfo, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(VolumeBackupInfo)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderConfig) DeepCopyInto(out *CloudProviderConfig) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSConfig)
			**out = **in
		}
	}
	if in.GCP != nil {
		in, out := &in.GCP, &out.GCP
		if *in == nil {
			*out = nil
		} else {
			*out = new(GCPConfig)
			**out = **in
		}
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		if *in == nil {
			*out = nil
		} else {
			*out = new(AzureConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderConfig.
func (in *CloudProviderConfig) DeepCopy() *CloudProviderConfig {
	if in == nil {
		return nil
	}
	out := new(CloudProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.PersistentVolumeProvider != nil {
		in, out := &in.PersistentVolumeProvider, &out.PersistentVolumeProvider
		if *in == nil {
			*out = nil
		} else {
			*out = new(CloudProviderConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	in.BackupStorageProvider.DeepCopyInto(&out.BackupStorageProvider)
	out.BackupSyncPeriod = in.BackupSyncPeriod
	out.GCSyncPeriod = in.GCSyncPeriod
	out.ScheduleSyncPeriod = in.ScheduleSyncPeriod
	if in.ResourcePriorities != nil {
		in, out := &in.ResourcePriorities, &out.ResourcePriorities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigList) DeepCopyInto(out *ConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigList.
func (in *ConfigList) DeepCopy() *ConfigList {
	if in == nil {
		return nil
	}
	out := new(ConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloadRequest) DeepCopyInto(out *DownloadRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloadRequest.
func (in *DownloadRequest) DeepCopy() *DownloadRequest {
	if in == nil {
		return nil
	}
	out := new(DownloadRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DownloadRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloadRequestList) DeepCopyInto(out *DownloadRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DownloadRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloadRequestList.
func (in *DownloadRequestList) DeepCopy() *DownloadRequestList {
	if in == nil {
		return nil
	}
	out := new(DownloadRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DownloadRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloadRequestSpec) DeepCopyInto(out *DownloadRequestSpec) {
	*out = *in
	out.Target = in.Target
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloadRequestSpec.
func (in *DownloadRequestSpec) DeepCopy() *DownloadRequestSpec {
	if in == nil {
		return nil
	}
	out := new(DownloadRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloadRequestStatus) DeepCopyInto(out *DownloadRequestStatus) {
	*out = *in
	in.Expiration.DeepCopyInto(&out.Expiration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloadRequestStatus.
func (in *DownloadRequestStatus) DeepCopy() *DownloadRequestStatus {
	if in == nil {
		return nil
	}
	out := new(DownloadRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloadTarget) DeepCopyInto(out *DownloadTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloadTarget.
func (in *DownloadTarget) DeepCopy() *DownloadTarget {
	if in == nil {
		return nil
	}
	out := new(DownloadTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecHook) DeepCopyInto(out *ExecHook) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecHook.
func (in *ExecHook) DeepCopy() *ExecHook {
	if in == nil {
		return nil
	}
	out := new(ExecHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPConfig) DeepCopyInto(out *GCPConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPConfig.
func (in *GCPConfig) DeepCopy() *GCPConfig {
	if in == nil {
		return nil
	}
	out := new(GCPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageProviderConfig) DeepCopyInto(out *ObjectStorageProviderConfig) {
	*out = *in
	in.CloudProviderConfig.DeepCopyInto(&out.CloudProviderConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageProviderConfig.
func (in *ObjectStorageProviderConfig) DeepCopy() *ObjectStorageProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restore) DeepCopyInto(out *Restore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restore.
func (in *Restore) DeepCopy() *Restore {
	if in == nil {
		return nil
	}
	out := new(Restore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Restore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreList) DeepCopyInto(out *RestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Restore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreList.
func (in *RestoreList) DeepCopy() *RestoreList {
	if in == nil {
		return nil
	}
	out := new(RestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreResult) DeepCopyInto(out *RestoreResult) {
	*out = *in
	if in.Ark != nil {
		in, out := &in.Ark, &out.Ark
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreResult.
func (in *RestoreResult) DeepCopy() *RestoreResult {
	if in == nil {
		return nil
	}
	out := new(RestoreResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreSpec) DeepCopyInto(out *RestoreSpec) {
	*out = *in
	if in.IncludedNamespaces != nil {
		in, out := &in.IncludedNamespaces, &out.IncludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedNamespaces != nil {
		in, out := &in.ExcludedNamespaces, &out.ExcludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludedResources != nil {
		in, out := &in.IncludedResources, &out.IncludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceMapping != nil {
		in, out := &in.NamespaceMapping, &out.NamespaceMapping
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RestorePVs != nil {
		in, out := &in.RestorePVs, &out.RestorePVs
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.IncludeClusterResources != nil {
		in, out := &in.IncludeClusterResources, &out.IncludeClusterResources
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreSpec.
func (in *RestoreSpec) DeepCopy() *RestoreSpec {
	if in == nil {
		return nil
	}
	out := new(RestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreStatus) DeepCopyInto(out *RestoreStatus) {
	*out = *in
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreStatus.
func (in *RestoreStatus) DeepCopy() *RestoreStatus {
	if in == nil {
		return nil
	}
	out := new(RestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schedule) DeepCopyInto(out *Schedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schedule.
func (in *Schedule) DeepCopy() *Schedule {
	if in == nil {
		return nil
	}
	out := new(Schedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Schedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleList) DeepCopyInto(out *ScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Schedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleList.
func (in *ScheduleList) DeepCopy() *ScheduleList {
	if in == nil {
		return nil
	}
	out := new(ScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleSpec) DeepCopyInto(out *ScheduleSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleSpec.
func (in *ScheduleSpec) DeepCopy() *ScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleStatus) DeepCopyInto(out *ScheduleStatus) {
	*out = *in
	in.LastBackup.DeepCopyInto(&out.LastBackup)
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleStatus.
func (in *ScheduleStatus) DeepCopy() *ScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupInfo) DeepCopyInto(out *VolumeBackupInfo) {
	*out = *in
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupInfo.
func (in *VolumeBackupInfo) DeepCopy() *VolumeBackupInfo {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupInfo)
	in.DeepCopyInto(out)
	return out
}
