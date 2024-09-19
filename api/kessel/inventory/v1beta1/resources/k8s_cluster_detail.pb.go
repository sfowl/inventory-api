// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/resources/k8s_cluster_detail.proto

package resources

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// the aggregate status of the cluster
type K8SClusterDetail_ClusterStatus int32

const (
	K8SClusterDetail_CLUSTER_STATUS_UNSPECIFIED K8SClusterDetail_ClusterStatus = 0
	K8SClusterDetail_CLUSTER_STATUS_OTHER       K8SClusterDetail_ClusterStatus = 1
	K8SClusterDetail_READY                      K8SClusterDetail_ClusterStatus = 2
	K8SClusterDetail_FAILED                     K8SClusterDetail_ClusterStatus = 3
	K8SClusterDetail_OFFLINE                    K8SClusterDetail_ClusterStatus = 4
)

// Enum value maps for K8SClusterDetail_ClusterStatus.
var (
	K8SClusterDetail_ClusterStatus_name = map[int32]string{
		0: "CLUSTER_STATUS_UNSPECIFIED",
		1: "CLUSTER_STATUS_OTHER",
		2: "READY",
		3: "FAILED",
		4: "OFFLINE",
	}
	K8SClusterDetail_ClusterStatus_value = map[string]int32{
		"CLUSTER_STATUS_UNSPECIFIED": 0,
		"CLUSTER_STATUS_OTHER":       1,
		"READY":                      2,
		"FAILED":                     3,
		"OFFLINE":                    4,
	}
)

func (x K8SClusterDetail_ClusterStatus) Enum() *K8SClusterDetail_ClusterStatus {
	p := new(K8SClusterDetail_ClusterStatus)
	*p = x
	return p
}

func (x K8SClusterDetail_ClusterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SClusterDetail_ClusterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes[0].Descriptor()
}

func (K8SClusterDetail_ClusterStatus) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes[0]
}

func (x K8SClusterDetail_ClusterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SClusterDetail_ClusterStatus.Descriptor instead.
func (K8SClusterDetail_ClusterStatus) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescGZIP(), []int{0, 0}
}

// The kubernetes vendor
type K8SClusterDetail_KubeVendor int32

const (
	K8SClusterDetail_KUBE_VENDOR_UNSPECIFIED K8SClusterDetail_KubeVendor = 0
	K8SClusterDetail_KUBE_VENDOR_OTHER       K8SClusterDetail_KubeVendor = 1
	K8SClusterDetail_AKS                     K8SClusterDetail_KubeVendor = 2
	K8SClusterDetail_EKS                     K8SClusterDetail_KubeVendor = 3
	K8SClusterDetail_IKS                     K8SClusterDetail_KubeVendor = 4
	K8SClusterDetail_OPENSHIFT               K8SClusterDetail_KubeVendor = 5
	K8SClusterDetail_GKE                     K8SClusterDetail_KubeVendor = 6
)

// Enum value maps for K8SClusterDetail_KubeVendor.
var (
	K8SClusterDetail_KubeVendor_name = map[int32]string{
		0: "KUBE_VENDOR_UNSPECIFIED",
		1: "KUBE_VENDOR_OTHER",
		2: "AKS",
		3: "EKS",
		4: "IKS",
		5: "OPENSHIFT",
		6: "GKE",
	}
	K8SClusterDetail_KubeVendor_value = map[string]int32{
		"KUBE_VENDOR_UNSPECIFIED": 0,
		"KUBE_VENDOR_OTHER":       1,
		"AKS":                     2,
		"EKS":                     3,
		"IKS":                     4,
		"OPENSHIFT":               5,
		"GKE":                     6,
	}
)

func (x K8SClusterDetail_KubeVendor) Enum() *K8SClusterDetail_KubeVendor {
	p := new(K8SClusterDetail_KubeVendor)
	*p = x
	return p
}

func (x K8SClusterDetail_KubeVendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SClusterDetail_KubeVendor) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes[1].Descriptor()
}

func (K8SClusterDetail_KubeVendor) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes[1]
}

func (x K8SClusterDetail_KubeVendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SClusterDetail_KubeVendor.Descriptor instead.
func (K8SClusterDetail_KubeVendor) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescGZIP(), []int{0, 1}
}

// The platform on which this cluster is hosted
type K8SClusterDetail_CloudPlatform int32

const (
	K8SClusterDetail_CLOUD_PLATFORM_UNSPECIFIED K8SClusterDetail_CloudPlatform = 0
	K8SClusterDetail_CLOUD_PLATFORM_OTHER       K8SClusterDetail_CloudPlatform = 1
	K8SClusterDetail_NONE_UPI                   K8SClusterDetail_CloudPlatform = 2
	K8SClusterDetail_BAREMETAL_IPI              K8SClusterDetail_CloudPlatform = 3
	K8SClusterDetail_BAREMETAL_UPI              K8SClusterDetail_CloudPlatform = 4
	K8SClusterDetail_AWS_IPI                    K8SClusterDetail_CloudPlatform = 5
	K8SClusterDetail_AWS_UPI                    K8SClusterDetail_CloudPlatform = 6
	K8SClusterDetail_AZURE_IPI                  K8SClusterDetail_CloudPlatform = 7
	K8SClusterDetail_AZURE_UPI                  K8SClusterDetail_CloudPlatform = 8
	K8SClusterDetail_IBMCLOUD_IPI               K8SClusterDetail_CloudPlatform = 9
	K8SClusterDetail_IBMCLOUD_UPI               K8SClusterDetail_CloudPlatform = 10
	K8SClusterDetail_KUBEVIRT_IPI               K8SClusterDetail_CloudPlatform = 11
	K8SClusterDetail_OPENSTACK_IPI              K8SClusterDetail_CloudPlatform = 12
	K8SClusterDetail_OPENSTACK_UPI              K8SClusterDetail_CloudPlatform = 13
	K8SClusterDetail_GCP_IPI                    K8SClusterDetail_CloudPlatform = 14
	K8SClusterDetail_GCP_UPI                    K8SClusterDetail_CloudPlatform = 15
	K8SClusterDetail_NUTANIX_IPI                K8SClusterDetail_CloudPlatform = 16
	K8SClusterDetail_NUTANIX_UPI                K8SClusterDetail_CloudPlatform = 17
	K8SClusterDetail_VSPHERE_IPI                K8SClusterDetail_CloudPlatform = 18
	K8SClusterDetail_VSPHERE_UPI                K8SClusterDetail_CloudPlatform = 19
	K8SClusterDetail_OVIRT_IPI                  K8SClusterDetail_CloudPlatform = 20
)

// Enum value maps for K8SClusterDetail_CloudPlatform.
var (
	K8SClusterDetail_CloudPlatform_name = map[int32]string{
		0:  "CLOUD_PLATFORM_UNSPECIFIED",
		1:  "CLOUD_PLATFORM_OTHER",
		2:  "NONE_UPI",
		3:  "BAREMETAL_IPI",
		4:  "BAREMETAL_UPI",
		5:  "AWS_IPI",
		6:  "AWS_UPI",
		7:  "AZURE_IPI",
		8:  "AZURE_UPI",
		9:  "IBMCLOUD_IPI",
		10: "IBMCLOUD_UPI",
		11: "KUBEVIRT_IPI",
		12: "OPENSTACK_IPI",
		13: "OPENSTACK_UPI",
		14: "GCP_IPI",
		15: "GCP_UPI",
		16: "NUTANIX_IPI",
		17: "NUTANIX_UPI",
		18: "VSPHERE_IPI",
		19: "VSPHERE_UPI",
		20: "OVIRT_IPI",
	}
	K8SClusterDetail_CloudPlatform_value = map[string]int32{
		"CLOUD_PLATFORM_UNSPECIFIED": 0,
		"CLOUD_PLATFORM_OTHER":       1,
		"NONE_UPI":                   2,
		"BAREMETAL_IPI":              3,
		"BAREMETAL_UPI":              4,
		"AWS_IPI":                    5,
		"AWS_UPI":                    6,
		"AZURE_IPI":                  7,
		"AZURE_UPI":                  8,
		"IBMCLOUD_IPI":               9,
		"IBMCLOUD_UPI":               10,
		"KUBEVIRT_IPI":               11,
		"OPENSTACK_IPI":              12,
		"OPENSTACK_UPI":              13,
		"GCP_IPI":                    14,
		"GCP_UPI":                    15,
		"NUTANIX_IPI":                16,
		"NUTANIX_UPI":                17,
		"VSPHERE_IPI":                18,
		"VSPHERE_UPI":                19,
		"OVIRT_IPI":                  20,
	}
)

func (x K8SClusterDetail_CloudPlatform) Enum() *K8SClusterDetail_CloudPlatform {
	p := new(K8SClusterDetail_CloudPlatform)
	*p = x
	return p
}

func (x K8SClusterDetail_CloudPlatform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SClusterDetail_CloudPlatform) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes[2].Descriptor()
}

func (K8SClusterDetail_CloudPlatform) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes[2]
}

func (x K8SClusterDetail_CloudPlatform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SClusterDetail_CloudPlatform.Descriptor instead.
func (K8SClusterDetail_CloudPlatform) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescGZIP(), []int{0, 2}
}

type K8SClusterDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The OCP cluster ID or ARN etc for *KS
	ExternalClusterId string                         `protobuf:"bytes,219571597,opt,name=external_cluster_id,proto3" json:"external_cluster_id,omitempty"`
	ClusterStatus     K8SClusterDetail_ClusterStatus `protobuf:"varint,499346904,opt,name=cluster_status,proto3,enum=kessel.inventory.v1beta1.resources.K8SClusterDetail_ClusterStatus" json:"cluster_status,omitempty"`
	// The version of kubernetes
	KubeVersion string                      `protobuf:"bytes,395858490,opt,name=kube_version,proto3" json:"kube_version,omitempty"`
	KubeVendor  K8SClusterDetail_KubeVendor `protobuf:"varint,264191642,opt,name=kube_vendor,proto3,enum=kessel.inventory.v1beta1.resources.K8SClusterDetail_KubeVendor" json:"kube_vendor,omitempty"`
	// The version of the productized kubernetes distribution
	VendorVersion string                         `protobuf:"bytes,23961827,opt,name=vendor_version,proto3" json:"vendor_version,omitempty"`
	CloudPlatform K8SClusterDetail_CloudPlatform `protobuf:"varint,476768062,opt,name=cloud_platform,proto3,enum=kessel.inventory.v1beta1.resources.K8SClusterDetail_CloudPlatform" json:"cloud_platform,omitempty"`
	Nodes         []*K8SClusterDetailNodesInner  `protobuf:"bytes,75440785,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *K8SClusterDetail) Reset() {
	*x = K8SClusterDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SClusterDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SClusterDetail) ProtoMessage() {}

func (x *K8SClusterDetail) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SClusterDetail.ProtoReflect.Descriptor instead.
func (*K8SClusterDetail) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescGZIP(), []int{0}
}

func (x *K8SClusterDetail) GetExternalClusterId() string {
	if x != nil {
		return x.ExternalClusterId
	}
	return ""
}

func (x *K8SClusterDetail) GetClusterStatus() K8SClusterDetail_ClusterStatus {
	if x != nil {
		return x.ClusterStatus
	}
	return K8SClusterDetail_CLUSTER_STATUS_UNSPECIFIED
}

func (x *K8SClusterDetail) GetKubeVersion() string {
	if x != nil {
		return x.KubeVersion
	}
	return ""
}

func (x *K8SClusterDetail) GetKubeVendor() K8SClusterDetail_KubeVendor {
	if x != nil {
		return x.KubeVendor
	}
	return K8SClusterDetail_KUBE_VENDOR_UNSPECIFIED
}

func (x *K8SClusterDetail) GetVendorVersion() string {
	if x != nil {
		return x.VendorVersion
	}
	return ""
}

func (x *K8SClusterDetail) GetCloudPlatform() K8SClusterDetail_CloudPlatform {
	if x != nil {
		return x.CloudPlatform
	}
	return K8SClusterDetail_CLOUD_PLATFORM_UNSPECIFIED
}

func (x *K8SClusterDetail) GetNodes() []*K8SClusterDetailNodesInner {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6b,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0x47, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x09, 0x0a, 0x10, 0x4b, 0x38, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x13,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x8d, 0xcb, 0xd9, 0x68, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x7b, 0x0a, 0x0e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0xd8, 0xdb, 0x8d,
	0xee, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x82,
	0x01, 0x05, 0x10, 0x01, 0x22, 0x01, 0x00, 0x52, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0c, 0x6b, 0x75, 0x62, 0x65, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xba, 0xa4, 0xe1, 0xbc, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x71, 0x0a, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x9a,
	0xfd, 0xfc, 0x7d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x4b, 0x75,
	0x62, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x82, 0x01, 0x05,
	0x10, 0x01, 0x22, 0x01, 0x00, 0x52, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x12, 0x32, 0x0a, 0x0e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0xe3, 0xc1, 0xb6, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x7b, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0xbe, 0xce, 0xab, 0xe3, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x42, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x82, 0x01, 0x05, 0x10, 0x01,
	0x22, 0x01, 0x00, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x66, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x91, 0xc5, 0xfc,
	0x23, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0x92, 0x01, 0x07, 0x08, 0x01, 0x22,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x0d, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f,
	0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x04, 0x22, 0x73, 0x0a, 0x0a, 0x4b, 0x75,
	0x62, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x17, 0x4b, 0x55, 0x42, 0x45,
	0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4b, 0x55, 0x42, 0x45, 0x5f, 0x56, 0x45,
	0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x4b, 0x53, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4b, 0x53, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x49, 0x4b, 0x53, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x50, 0x45, 0x4e, 0x53,
	0x48, 0x49, 0x46, 0x54, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x4b, 0x45, 0x10, 0x06, 0x22,
	0xfe, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4e,
	0x4f, 0x4e, 0x45, 0x5f, 0x55, 0x50, 0x49, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x41, 0x52,
	0x45, 0x4d, 0x45, 0x54, 0x41, 0x4c, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x42, 0x41, 0x52, 0x45, 0x4d, 0x45, 0x54, 0x41, 0x4c, 0x5f, 0x55, 0x50, 0x49, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x57, 0x53, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x57, 0x53, 0x5f, 0x55, 0x50, 0x49, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x5a, 0x55,
	0x52, 0x45, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x5a, 0x55, 0x52,
	0x45, 0x5f, 0x55, 0x50, 0x49, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x42, 0x4d, 0x43, 0x4c,
	0x4f, 0x55, 0x44, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x42, 0x4d,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x55, 0x50, 0x49, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4b,
	0x55, 0x42, 0x45, 0x56, 0x49, 0x52, 0x54, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x0b, 0x12, 0x11, 0x0a,
	0x0d, 0x4f, 0x50, 0x45, 0x4e, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x0c,
	0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x45, 0x4e, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x55, 0x50,
	0x49, 0x10, 0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x43, 0x50, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x0e,
	0x12, 0x0b, 0x0a, 0x07, 0x47, 0x43, 0x50, 0x5f, 0x55, 0x50, 0x49, 0x10, 0x0f, 0x12, 0x0f, 0x0a,
	0x0b, 0x4e, 0x55, 0x54, 0x41, 0x4e, 0x49, 0x58, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x10, 0x12, 0x0f,
	0x0a, 0x0b, 0x4e, 0x55, 0x54, 0x41, 0x4e, 0x49, 0x58, 0x5f, 0x55, 0x50, 0x49, 0x10, 0x11, 0x12,
	0x0f, 0x0a, 0x0b, 0x56, 0x53, 0x50, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x12,
	0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x53, 0x50, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x55, 0x50, 0x49, 0x10,
	0x13, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x56, 0x49, 0x52, 0x54, 0x5f, 0x49, 0x50, 0x49, 0x10, 0x14,
	0x42, 0x9d, 0x01, 0x0a, 0x32, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x15, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescData = file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDesc
)

func file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDescData
}

var file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_goTypes = []any{
	(K8SClusterDetail_ClusterStatus)(0), // 0: kessel.inventory.v1beta1.resources.K8sClusterDetail.ClusterStatus
	(K8SClusterDetail_KubeVendor)(0),    // 1: kessel.inventory.v1beta1.resources.K8sClusterDetail.KubeVendor
	(K8SClusterDetail_CloudPlatform)(0), // 2: kessel.inventory.v1beta1.resources.K8sClusterDetail.CloudPlatform
	(*K8SClusterDetail)(nil),            // 3: kessel.inventory.v1beta1.resources.K8sClusterDetail
	(*K8SClusterDetailNodesInner)(nil),  // 4: kessel.inventory.v1beta1.resources.K8sClusterDetailNodesInner
}
var file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_depIdxs = []int32{
	0, // 0: kessel.inventory.v1beta1.resources.K8sClusterDetail.cluster_status:type_name -> kessel.inventory.v1beta1.resources.K8sClusterDetail.ClusterStatus
	1, // 1: kessel.inventory.v1beta1.resources.K8sClusterDetail.kube_vendor:type_name -> kessel.inventory.v1beta1.resources.K8sClusterDetail.KubeVendor
	2, // 2: kessel.inventory.v1beta1.resources.K8sClusterDetail.cloud_platform:type_name -> kessel.inventory.v1beta1.resources.K8sClusterDetail.CloudPlatform
	4, // 3: kessel.inventory.v1beta1.resources.K8sClusterDetail.nodes:type_name -> kessel.inventory.v1beta1.resources.K8sClusterDetailNodesInner
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_init() }
func file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_init() {
	if File_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_nodes_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*K8SClusterDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_depIdxs,
		EnumInfos:         file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_enumTypes,
		MessageInfos:      file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto = out.File
	file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_goTypes = nil
	file_kessel_inventory_v1beta1_resources_k8s_cluster_detail_proto_depIdxs = nil
}
