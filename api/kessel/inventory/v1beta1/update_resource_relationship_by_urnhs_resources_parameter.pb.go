// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/update_resource_relationship_by_urnhs_resources_parameter.proto

package v1beta1

import (
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

type UpdateResourceRelationshipByUrnHsResourcesParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectResource string `protobuf:"bytes,19549760,opt,name=subject_resource,json=subjectResource,proto3" json:"subject_resource,omitempty"`
	ObjectResource  string `protobuf:"bytes,254445841,opt,name=object_resource,json=objectResource,proto3" json:"object_resource,omitempty"`
}

func (x *UpdateResourceRelationshipByUrnHsResourcesParameter) Reset() {
	*x = UpdateResourceRelationshipByUrnHsResourcesParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceRelationshipByUrnHsResourcesParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceRelationshipByUrnHsResourcesParameter) ProtoMessage() {}

func (x *UpdateResourceRelationshipByUrnHsResourcesParameter) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceRelationshipByUrnHsResourcesParameter.ProtoReflect.Descriptor instead.
func (*UpdateResourceRelationshipByUrnHsResourcesParameter) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateResourceRelationshipByUrnHsResourcesParameter) GetSubjectResource() string {
	if x != nil {
		return x.SubjectResource
	}
	return ""
}

func (x *UpdateResourceRelationshipByUrnHsResourcesParameter) GetObjectResource() string {
	if x != nil {
		return x.ObjectResource
	}
	return ""
}

var File_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDesc = []byte{
	0x0a, 0x58, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x72, 0x6e, 0x68, 0x73,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6b, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x22, 0x8f, 0x01, 0x0a, 0x33, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e, 0x48, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x10,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0xc0, 0x9c, 0xa9, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x91, 0x92,
	0xaa, 0x79, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x66, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescData = file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDesc
)

func file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDescData
}

var file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_goTypes = []any{
	(*UpdateResourceRelationshipByUrnHsResourcesParameter)(nil), // 0: kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsResourcesParameter
}
var file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_init()
}
func file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_init() {
	if File_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateResourceRelationshipByUrnHsResourcesParameter); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto = out.File
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_goTypes = nil
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_depIdxs = nil
}
