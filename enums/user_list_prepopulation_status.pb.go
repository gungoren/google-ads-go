// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.2
// source: google/ads/googleads/v8/enums/user_list_prepopulation_status.proto

package enums

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible user list prepopulation status.
type UserListPrepopulationStatusEnum_UserListPrepopulationStatus int32

const (
	// Not specified.
	UserListPrepopulationStatusEnum_UNSPECIFIED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 0
	// Used for return value only. Represents value unknown in this version.
	UserListPrepopulationStatusEnum_UNKNOWN UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 1
	// Prepopoulation is being requested.
	UserListPrepopulationStatusEnum_REQUESTED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 2
	// Prepopulation is finished.
	UserListPrepopulationStatusEnum_FINISHED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 3
	// Prepopulation failed.
	UserListPrepopulationStatusEnum_FAILED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 4
)

// Enum value maps for UserListPrepopulationStatusEnum_UserListPrepopulationStatus.
var (
	UserListPrepopulationStatusEnum_UserListPrepopulationStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "REQUESTED",
		3: "FINISHED",
		4: "FAILED",
	}
	UserListPrepopulationStatusEnum_UserListPrepopulationStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"REQUESTED":   2,
		"FINISHED":    3,
		"FAILED":      4,
	}
)

func (x UserListPrepopulationStatusEnum_UserListPrepopulationStatus) Enum() *UserListPrepopulationStatusEnum_UserListPrepopulationStatus {
	p := new(UserListPrepopulationStatusEnum_UserListPrepopulationStatus)
	*p = x
	return p
}

func (x UserListPrepopulationStatusEnum_UserListPrepopulationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserListPrepopulationStatusEnum_UserListPrepopulationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_user_list_prepopulation_status_proto_enumTypes[0].Descriptor()
}

func (UserListPrepopulationStatusEnum_UserListPrepopulationStatus) Type() protoreflect.EnumType {
	return &file_enums_user_list_prepopulation_status_proto_enumTypes[0]
}

func (x UserListPrepopulationStatusEnum_UserListPrepopulationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserListPrepopulationStatusEnum_UserListPrepopulationStatus.Descriptor instead.
func (UserListPrepopulationStatusEnum_UserListPrepopulationStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_user_list_prepopulation_status_proto_rawDescGZIP(), []int{0, 0}
}

// Indicates status of prepopulation based on the rule.
type UserListPrepopulationStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserListPrepopulationStatusEnum) Reset() {
	*x = UserListPrepopulationStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_user_list_prepopulation_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListPrepopulationStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListPrepopulationStatusEnum) ProtoMessage() {}

func (x *UserListPrepopulationStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_user_list_prepopulation_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListPrepopulationStatusEnum.ProtoReflect.Descriptor instead.
func (*UserListPrepopulationStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_user_list_prepopulation_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_user_list_prepopulation_status_proto protoreflect.FileDescriptor

var file_enums_user_list_prepopulation_status_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x70, 0x6f, 0x70,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x87, 0x01, 0x0a, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x65, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x64, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x65, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0xf5, 0x01, 0x0a, 0x21,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x20, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x70, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_user_list_prepopulation_status_proto_rawDescOnce sync.Once
	file_enums_user_list_prepopulation_status_proto_rawDescData = file_enums_user_list_prepopulation_status_proto_rawDesc
)

func file_enums_user_list_prepopulation_status_proto_rawDescGZIP() []byte {
	file_enums_user_list_prepopulation_status_proto_rawDescOnce.Do(func() {
		file_enums_user_list_prepopulation_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_user_list_prepopulation_status_proto_rawDescData)
	})
	return file_enums_user_list_prepopulation_status_proto_rawDescData
}

var file_enums_user_list_prepopulation_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_user_list_prepopulation_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_user_list_prepopulation_status_proto_goTypes = []interface{}{
	(UserListPrepopulationStatusEnum_UserListPrepopulationStatus)(0), // 0: google.ads.googleads.v8.enums.UserListPrepopulationStatusEnum.UserListPrepopulationStatus
	(*UserListPrepopulationStatusEnum)(nil),                          // 1: google.ads.googleads.v8.enums.UserListPrepopulationStatusEnum
}
var file_enums_user_list_prepopulation_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_user_list_prepopulation_status_proto_init() }
func file_enums_user_list_prepopulation_status_proto_init() {
	if File_enums_user_list_prepopulation_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_user_list_prepopulation_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListPrepopulationStatusEnum); i {
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
			RawDescriptor: file_enums_user_list_prepopulation_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_user_list_prepopulation_status_proto_goTypes,
		DependencyIndexes: file_enums_user_list_prepopulation_status_proto_depIdxs,
		EnumInfos:         file_enums_user_list_prepopulation_status_proto_enumTypes,
		MessageInfos:      file_enums_user_list_prepopulation_status_proto_msgTypes,
	}.Build()
	File_enums_user_list_prepopulation_status_proto = out.File
	file_enums_user_list_prepopulation_status_proto_rawDesc = nil
	file_enums_user_list_prepopulation_status_proto_goTypes = nil
	file_enums_user_list_prepopulation_status_proto_depIdxs = nil
}
