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
// source: google/ads/googleads/v8/enums/ad_destination_type.proto

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

// Enumerates Google Ads destination types
type AdDestinationTypeEnum_AdDestinationType int32

const (
	// Not specified.
	AdDestinationTypeEnum_UNSPECIFIED AdDestinationTypeEnum_AdDestinationType = 0
	// The value is unknown in this version.
	AdDestinationTypeEnum_UNKNOWN AdDestinationTypeEnum_AdDestinationType = 1
	// Ads that don't intend to drive users off from ads to other destinations
	AdDestinationTypeEnum_NOT_APPLICABLE AdDestinationTypeEnum_AdDestinationType = 2
	// Website
	AdDestinationTypeEnum_WEBSITE AdDestinationTypeEnum_AdDestinationType = 3
	// App Deep Link
	AdDestinationTypeEnum_APP_DEEP_LINK AdDestinationTypeEnum_AdDestinationType = 4
	// iOS App Store or Play Store
	AdDestinationTypeEnum_APP_STORE AdDestinationTypeEnum_AdDestinationType = 5
	// Call Dialer
	AdDestinationTypeEnum_PHONE_CALL AdDestinationTypeEnum_AdDestinationType = 6
	// Map App
	AdDestinationTypeEnum_MAP_DIRECTIONS AdDestinationTypeEnum_AdDestinationType = 7
	// Location Dedicated Page
	AdDestinationTypeEnum_LOCATION_LISTING AdDestinationTypeEnum_AdDestinationType = 8
	// Text Message
	AdDestinationTypeEnum_MESSAGE AdDestinationTypeEnum_AdDestinationType = 9
	// Lead Generation Form
	AdDestinationTypeEnum_LEAD_FORM AdDestinationTypeEnum_AdDestinationType = 10
	// YouTube
	AdDestinationTypeEnum_YOUTUBE AdDestinationTypeEnum_AdDestinationType = 11
	// Ad Destination for Conversions with keys unknown
	AdDestinationTypeEnum_UNMODELED_FOR_CONVERSIONS AdDestinationTypeEnum_AdDestinationType = 12
)

// Enum value maps for AdDestinationTypeEnum_AdDestinationType.
var (
	AdDestinationTypeEnum_AdDestinationType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "NOT_APPLICABLE",
		3:  "WEBSITE",
		4:  "APP_DEEP_LINK",
		5:  "APP_STORE",
		6:  "PHONE_CALL",
		7:  "MAP_DIRECTIONS",
		8:  "LOCATION_LISTING",
		9:  "MESSAGE",
		10: "LEAD_FORM",
		11: "YOUTUBE",
		12: "UNMODELED_FOR_CONVERSIONS",
	}
	AdDestinationTypeEnum_AdDestinationType_value = map[string]int32{
		"UNSPECIFIED":               0,
		"UNKNOWN":                   1,
		"NOT_APPLICABLE":            2,
		"WEBSITE":                   3,
		"APP_DEEP_LINK":             4,
		"APP_STORE":                 5,
		"PHONE_CALL":                6,
		"MAP_DIRECTIONS":            7,
		"LOCATION_LISTING":          8,
		"MESSAGE":                   9,
		"LEAD_FORM":                 10,
		"YOUTUBE":                   11,
		"UNMODELED_FOR_CONVERSIONS": 12,
	}
)

func (x AdDestinationTypeEnum_AdDestinationType) Enum() *AdDestinationTypeEnum_AdDestinationType {
	p := new(AdDestinationTypeEnum_AdDestinationType)
	*p = x
	return p
}

func (x AdDestinationTypeEnum_AdDestinationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdDestinationTypeEnum_AdDestinationType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_ad_destination_type_proto_enumTypes[0].Descriptor()
}

func (AdDestinationTypeEnum_AdDestinationType) Type() protoreflect.EnumType {
	return &file_enums_ad_destination_type_proto_enumTypes[0]
}

func (x AdDestinationTypeEnum_AdDestinationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdDestinationTypeEnum_AdDestinationType.Descriptor instead.
func (AdDestinationTypeEnum_AdDestinationType) EnumDescriptor() ([]byte, []int) {
	return file_enums_ad_destination_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enumeration of Google Ads destination types.
type AdDestinationTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdDestinationTypeEnum) Reset() {
	*x = AdDestinationTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_ad_destination_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdDestinationTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdDestinationTypeEnum) ProtoMessage() {}

func (x *AdDestinationTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_ad_destination_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdDestinationTypeEnum.ProtoReflect.Descriptor instead.
func (*AdDestinationTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_ad_destination_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_ad_destination_type_proto protoreflect.FileDescriptor

var file_enums_ad_destination_type_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x61, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x15, 0x41, 0x64, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0xf6, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x4c,
	0x49, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x45, 0x42, 0x53,
	0x49, 0x54, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x50, 0x5f, 0x44, 0x45, 0x45,
	0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x50, 0x50, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x48, 0x4f, 0x4e, 0x45,
	0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x41, 0x50, 0x5f, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x08, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x09, 0x12, 0x0d,
	0x0a, 0x09, 0x4c, 0x45, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0x0a, 0x12, 0x0b, 0x0a,
	0x07, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x10, 0x0b, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x4e,
	0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x0c, 0x42, 0xeb, 0x01, 0x0a, 0x21, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42,
	0x16, 0x41, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_ad_destination_type_proto_rawDescOnce sync.Once
	file_enums_ad_destination_type_proto_rawDescData = file_enums_ad_destination_type_proto_rawDesc
)

func file_enums_ad_destination_type_proto_rawDescGZIP() []byte {
	file_enums_ad_destination_type_proto_rawDescOnce.Do(func() {
		file_enums_ad_destination_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_ad_destination_type_proto_rawDescData)
	})
	return file_enums_ad_destination_type_proto_rawDescData
}

var file_enums_ad_destination_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_ad_destination_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_ad_destination_type_proto_goTypes = []interface{}{
	(AdDestinationTypeEnum_AdDestinationType)(0), // 0: google.ads.googleads.v8.enums.AdDestinationTypeEnum.AdDestinationType
	(*AdDestinationTypeEnum)(nil),                // 1: google.ads.googleads.v8.enums.AdDestinationTypeEnum
}
var file_enums_ad_destination_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_ad_destination_type_proto_init() }
func file_enums_ad_destination_type_proto_init() {
	if File_enums_ad_destination_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_ad_destination_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdDestinationTypeEnum); i {
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
			RawDescriptor: file_enums_ad_destination_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_ad_destination_type_proto_goTypes,
		DependencyIndexes: file_enums_ad_destination_type_proto_depIdxs,
		EnumInfos:         file_enums_ad_destination_type_proto_enumTypes,
		MessageInfos:      file_enums_ad_destination_type_proto_msgTypes,
	}.Build()
	File_enums_ad_destination_type_proto = out.File
	file_enums_ad_destination_type_proto_rawDesc = nil
	file_enums_ad_destination_type_proto_goTypes = nil
	file_enums_ad_destination_type_proto_depIdxs = nil
}
