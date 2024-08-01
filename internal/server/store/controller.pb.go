// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: controller/storage/servers/store/v1/controller.proto

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
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

// Controller contains all fields related to a Controller resource
type Controller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Private ID of the resource
	PrivateId string `protobuf:"bytes,10,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty"`
	// Address for the controller
	Address string `protobuf:"bytes,20,opt,name=address,proto3" json:"address,omitempty"`
	// Description of the resource (optional)
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,30,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// First seen time from the RDBMS
	CreateTime *timestamp.Timestamp `protobuf:"bytes,40,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Last time there was an update
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,50,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Controller) Reset() {
	*x = Controller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_servers_store_v1_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller) ProtoMessage() {}

func (x *Controller) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_servers_store_v1_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller.ProtoReflect.Descriptor instead.
func (*Controller) Descriptor() ([]byte, []int) {
	return file_controller_storage_servers_store_v1_controller_proto_rawDescGZIP(), []int{0}
}

func (x *Controller) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *Controller) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Controller) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Controller) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Controller) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_controller_storage_servers_store_v1_controller_proto protoreflect.FileDescriptor

var file_controller_storage_servers_store_v1_controller_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_servers_store_v1_controller_proto_rawDescOnce sync.Once
	file_controller_storage_servers_store_v1_controller_proto_rawDescData = file_controller_storage_servers_store_v1_controller_proto_rawDesc
)

func file_controller_storage_servers_store_v1_controller_proto_rawDescGZIP() []byte {
	file_controller_storage_servers_store_v1_controller_proto_rawDescOnce.Do(func() {
		file_controller_storage_servers_store_v1_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_servers_store_v1_controller_proto_rawDescData)
	})
	return file_controller_storage_servers_store_v1_controller_proto_rawDescData
}

var file_controller_storage_servers_store_v1_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_storage_servers_store_v1_controller_proto_goTypes = []any{
	(*Controller)(nil),          // 0: controller.storage.servers.store.v1.Controller
	(*timestamp.Timestamp)(nil), // 1: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_servers_store_v1_controller_proto_depIdxs = []int32{
	1, // 0: controller.storage.servers.store.v1.Controller.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	1, // 1: controller.storage.servers.store.v1.Controller.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_storage_servers_store_v1_controller_proto_init() }
func file_controller_storage_servers_store_v1_controller_proto_init() {
	if File_controller_storage_servers_store_v1_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_servers_store_v1_controller_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Controller); i {
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
			RawDescriptor: file_controller_storage_servers_store_v1_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_servers_store_v1_controller_proto_goTypes,
		DependencyIndexes: file_controller_storage_servers_store_v1_controller_proto_depIdxs,
		MessageInfos:      file_controller_storage_servers_store_v1_controller_proto_msgTypes,
	}.Build()
	File_controller_storage_servers_store_v1_controller_proto = out.File
	file_controller_storage_servers_store_v1_controller_proto_rawDesc = nil
	file_controller_storage_servers_store_v1_controller_proto_goTypes = nil
	file_controller_storage_servers_store_v1_controller_proto_depIdxs = nil
}
