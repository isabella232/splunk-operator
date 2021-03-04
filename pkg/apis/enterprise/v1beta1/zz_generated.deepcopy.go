// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundlePushInfo) DeepCopyInto(out *BundlePushInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundlePushInfo.
func (in *BundlePushInfo) DeepCopy() *BundlePushInfo {
	if in == nil {
		return nil
	}
	out := new(BundlePushInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheManagerSpec) DeepCopyInto(out *CacheManagerSpec) {
	*out = *in
	out.IndexAndCacheManagerCommonSpec = in.IndexAndCacheManagerCommonSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheManagerSpec.
func (in *CacheManagerSpec) DeepCopy() *CacheManagerSpec {
	if in == nil {
		return nil
	}
	out := new(CacheManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMaster) DeepCopyInto(out *ClusterMaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMaster.
func (in *ClusterMaster) DeepCopy() *ClusterMaster {
	if in == nil {
		return nil
	}
	out := new(ClusterMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterMaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMasterList) DeepCopyInto(out *ClusterMasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterMaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMasterList.
func (in *ClusterMasterList) DeepCopy() *ClusterMasterList {
	if in == nil {
		return nil
	}
	out := new(ClusterMasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterMasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMasterSpec) DeepCopyInto(out *ClusterMasterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMasterSpec.
func (in *ClusterMasterSpec) DeepCopy() *ClusterMasterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMasterStatus) DeepCopyInto(out *ClusterMasterStatus) {
	*out = *in
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	out.BundlePushTracker = in.BundlePushTracker
	if in.ResourceRevMap != nil {
		in, out := &in.ResourceRevMap, &out.ResourceRevMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMasterStatus.
func (in *ClusterMasterStatus) DeepCopy() *ClusterMasterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterMasterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSplunkSpec) DeepCopyInto(out *CommonSplunkSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.EtcVolumeStorageConfig = in.EtcVolumeStorageConfig
	out.VarVolumeStorageConfig = in.VarVolumeStorageConfig
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LicenseMasterRef = in.LicenseMasterRef
	out.ClusterMasterRef = in.ClusterMasterRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSplunkSpec.
func (in *CommonSplunkSpec) DeepCopy() *CommonSplunkSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSplunkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexAndCacheManagerCommonSpec) DeepCopyInto(out *IndexAndCacheManagerCommonSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexAndCacheManagerCommonSpec.
func (in *IndexAndCacheManagerCommonSpec) DeepCopy() *IndexAndCacheManagerCommonSpec {
	if in == nil {
		return nil
	}
	out := new(IndexAndCacheManagerCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexAndGlobalCommonSpec) DeepCopyInto(out *IndexAndGlobalCommonSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexAndGlobalCommonSpec.
func (in *IndexAndGlobalCommonSpec) DeepCopy() *IndexAndGlobalCommonSpec {
	if in == nil {
		return nil
	}
	out := new(IndexAndGlobalCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexConfDefaultsSpec) DeepCopyInto(out *IndexConfDefaultsSpec) {
	*out = *in
	out.IndexAndGlobalCommonSpec = in.IndexAndGlobalCommonSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexConfDefaultsSpec.
func (in *IndexConfDefaultsSpec) DeepCopy() *IndexConfDefaultsSpec {
	if in == nil {
		return nil
	}
	out := new(IndexConfDefaultsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexSpec) DeepCopyInto(out *IndexSpec) {
	*out = *in
	out.IndexAndCacheManagerCommonSpec = in.IndexAndCacheManagerCommonSpec
	out.IndexAndGlobalCommonSpec = in.IndexAndGlobalCommonSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexSpec.
func (in *IndexSpec) DeepCopy() *IndexSpec {
	if in == nil {
		return nil
	}
	out := new(IndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerCluster) DeepCopyInto(out *IndexerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerCluster.
func (in *IndexerCluster) DeepCopy() *IndexerCluster {
	if in == nil {
		return nil
	}
	out := new(IndexerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterList) DeepCopyInto(out *IndexerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IndexerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterList.
func (in *IndexerClusterList) DeepCopy() *IndexerClusterList {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterMemberStatus) DeepCopyInto(out *IndexerClusterMemberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterMemberStatus.
func (in *IndexerClusterMemberStatus) DeepCopy() *IndexerClusterMemberStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterSpec) DeepCopyInto(out *IndexerClusterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterSpec.
func (in *IndexerClusterSpec) DeepCopy() *IndexerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterStatus) DeepCopyInto(out *IndexerClusterStatus) {
	*out = *in
	if in.IndexerSecretChanged != nil {
		in, out := &in.IndexerSecretChanged, &out.IndexerSecretChanged
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.IdxcPasswordChangedSecrets != nil {
		in, out := &in.IdxcPasswordChangedSecrets, &out.IdxcPasswordChangedSecrets
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]IndexerClusterMemberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterStatus.
func (in *IndexerClusterStatus) DeepCopy() *IndexerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMaster) DeepCopyInto(out *LicenseMaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMaster.
func (in *LicenseMaster) DeepCopy() *LicenseMaster {
	if in == nil {
		return nil
	}
	out := new(LicenseMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterList) DeepCopyInto(out *LicenseMasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseMaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterList.
func (in *LicenseMasterList) DeepCopy() *LicenseMasterList {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterSpec) DeepCopyInto(out *LicenseMasterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterSpec.
func (in *LicenseMasterSpec) DeepCopy() *LicenseMasterSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterStatus) DeepCopyInto(out *LicenseMasterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterStatus.
func (in *LicenseMasterStatus) DeepCopy() *LicenseMasterStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadCluster) DeepCopyInto(out *SearchHeadCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadCluster.
func (in *SearchHeadCluster) DeepCopy() *SearchHeadCluster {
	if in == nil {
		return nil
	}
	out := new(SearchHeadCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterList) DeepCopyInto(out *SearchHeadClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SearchHeadCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterList.
func (in *SearchHeadClusterList) DeepCopy() *SearchHeadClusterList {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterMemberStatus) DeepCopyInto(out *SearchHeadClusterMemberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterMemberStatus.
func (in *SearchHeadClusterMemberStatus) DeepCopy() *SearchHeadClusterMemberStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterSpec) DeepCopyInto(out *SearchHeadClusterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	//out.SparkRef = in.SparkRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterSpec.
func (in *SearchHeadClusterSpec) DeepCopy() *SearchHeadClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterStatus) DeepCopyInto(out *SearchHeadClusterStatus) {
	*out = *in
	if in.ShcSecretChanged != nil {
		in, out := &in.ShcSecretChanged, &out.ShcSecretChanged
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.AdminSecretChanged != nil {
		in, out := &in.AdminSecretChanged, &out.AdminSecretChanged
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.AdminPasswordChangedSecrets != nil {
		in, out := &in.AdminPasswordChangedSecrets, &out.AdminPasswordChangedSecrets
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]SearchHeadClusterMemberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterStatus.
func (in *SearchHeadClusterStatus) DeepCopy() *SearchHeadClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartStoreSpec) DeepCopyInto(out *SmartStoreSpec) {
	*out = *in
	if in.VolList != nil {
		in, out := &in.VolList, &out.VolList
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.IndexList != nil {
		in, out := &in.IndexList, &out.IndexList
		*out = make([]IndexSpec, len(*in))
		copy(*out, *in)
	}
	out.Defaults = in.Defaults
	out.CacheManagerConf = in.CacheManagerConf
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartStoreSpec.
func (in *SmartStoreSpec) DeepCopy() *SmartStoreSpec {
	if in == nil {
		return nil
	}
	out := new(SmartStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
//func (in *Spark) DeepCopyInto(out *Spark) {
//	*out = *in
//	out.TypeMeta = in.TypeMeta
//	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
//	in.Spec.DeepCopyInto(&out.Spec)
//	out.Status = in.Status
//	return
//}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spark.
//func (in *Spark) DeepCopy() *Spark {
//	if in == nil {
//		return nil
//	}
//	out := new(Spark)
//	in.DeepCopyInto(out)
//	return out
//}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
//func (in *Spark) DeepCopyObject() runtime.Object {
//	if c := in.DeepCopy(); c != nil {
//		return c
//	}
//	return nil
//}
//
//// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
//func (in *SparkList) DeepCopyInto(out *SparkList) {
//	*out = *in
//	out.TypeMeta = in.TypeMeta
//	in.ListMeta.DeepCopyInto(&out.ListMeta)
//	if in.Items != nil {
//		in, out := &in.Items, &out.Items
//		*out = make([]Spark, len(*in))
//		for i := range *in {
//			(*in)[i].DeepCopyInto(&(*out)[i])
//		}
//	}
//	return
//}
//
//// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkList.
//func (in *SparkList) DeepCopy() *SparkList {
//	if in == nil {
//		return nil
//	}
//	out := new(SparkList)
//	in.DeepCopyInto(out)
//	return out
//}
//
//// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
//func (in *SparkList) DeepCopyObject() runtime.Object {
//	if c := in.DeepCopy(); c != nil {
//		return c
//	}
//	return nil
//}
//
//// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
//func (in *SparkSpec) DeepCopyInto(out *SparkSpec) {
//	*out = *in
//	in.Spec.DeepCopyInto(&out.Spec)
//	return
//}
//
//// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkSpec.
//func (in *SparkSpec) DeepCopy() *SparkSpec {
//	if in == nil {
//		return nil
//	}
//	out := new(SparkSpec)
//	in.DeepCopyInto(out)
//	return out
//}
//
//// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
//func (in *SparkStatus) DeepCopyInto(out *SparkStatus) {
//	*out = *in
//	return
//}
//
//// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkStatus.
//func (in *SparkStatus) DeepCopy() *SparkStatus {
//	if in == nil {
//		return nil
//	}
//	out := new(SparkStatus)
//	in.DeepCopyInto(out)
//	return out
//}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Standalone) DeepCopyInto(out *Standalone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Standalone.
func (in *Standalone) DeepCopy() *Standalone {
	if in == nil {
		return nil
	}
	out := new(Standalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Standalone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneList) DeepCopyInto(out *StandaloneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Standalone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneList.
func (in *StandaloneList) DeepCopy() *StandaloneList {
	if in == nil {
		return nil
	}
	out := new(StandaloneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandaloneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneSpec) DeepCopyInto(out *StandaloneSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	//out.SparkRef = in.SparkRef
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneSpec.
func (in *StandaloneSpec) DeepCopy() *StandaloneSpec {
	if in == nil {
		return nil
	}
	out := new(StandaloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneStatus) DeepCopyInto(out *StandaloneStatus) {
	*out = *in
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	if in.ResourceRevMap != nil {
		in, out := &in.ResourceRevMap, &out.ResourceRevMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneStatus.
func (in *StandaloneStatus) DeepCopy() *StandaloneStatus {
	if in == nil {
		return nil
	}
	out := new(StandaloneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassSpec) DeepCopyInto(out *StorageClassSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassSpec.
func (in *StorageClassSpec) DeepCopy() *StorageClassSpec {
	if in == nil {
		return nil
	}
	out := new(StorageClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
