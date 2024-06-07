// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: data/GiftStruct.proto

package data

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

// 更新时间 2024-04-29
type GiftStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,16,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GiftStruct) Reset() {
	*x = GiftStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_GiftStruct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftStruct) ProtoMessage() {}

func (x *GiftStruct) ProtoReflect() protoreflect.Message {
	mi := &file_data_GiftStruct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftStruct.ProtoReflect.Descriptor instead.
func (*GiftStruct) Descriptor() ([]byte, []int) {
	return file_data_GiftStruct_proto_rawDescGZIP(), []int{0}
}

func (x *GiftStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_data_GiftStruct_proto protoreflect.FileDescriptor

var file_data_GiftStruct_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x47, 0x69, 0x66, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f,
	0x68, 0x61, 0x63, 0x6b, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x69, 0x66, 0x74, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x6b, 0x0a, 0x45, 0x63, 0x6f, 0x6f, 0x6c, 0x2e, 0x73,
	0x63, 0x78, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50,
	0x01, 0x5a, 0x20, 0x53, 0x75, 0x6e, 0x6e, 0x79, 0x2f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f,
	0x68, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_GiftStruct_proto_rawDescOnce sync.Once
	file_data_GiftStruct_proto_rawDescData = file_data_GiftStruct_proto_rawDesc
)

func file_data_GiftStruct_proto_rawDescGZIP() []byte {
	file_data_GiftStruct_proto_rawDescOnce.Do(func() {
		file_data_GiftStruct_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_GiftStruct_proto_rawDescData)
	})
	return file_data_GiftStruct_proto_rawDescData
}

var file_data_GiftStruct_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_GiftStruct_proto_goTypes = []interface{}{
	(*GiftStruct)(nil), // 0: tiktok_hack.GiftStruct
}
var file_data_GiftStruct_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_GiftStruct_proto_init() }
func file_data_GiftStruct_proto_init() {
	if File_data_GiftStruct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_GiftStruct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftStruct); i {
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
			RawDescriptor: file_data_GiftStruct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_GiftStruct_proto_goTypes,
		DependencyIndexes: file_data_GiftStruct_proto_depIdxs,
		MessageInfos:      file_data_GiftStruct_proto_msgTypes,
	}.Build()
	File_data_GiftStruct_proto = out.File
	file_data_GiftStruct_proto_rawDesc = nil
	file_data_GiftStruct_proto_goTypes = nil
	file_data_GiftStruct_proto_depIdxs = nil
}
