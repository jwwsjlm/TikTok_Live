// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: data/Text.proto

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

// 更新时间 2024-06-06
type Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 键，可能用于唯一标识文本或作为字典的键
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// 默认模式，可能指定文本的默认展示格式或样式
	DefaultPattern string `protobuf:"bytes,2,opt,name=default_pattern,json=defaultPattern,proto3" json:"default_pattern,omitempty"`
	// 默认格式，可能包含文本的详细格式信息
	DefaultFormat *TextFormat `protobuf:"bytes,3,opt,name=default_format,json=defaultFormat,proto3" json:"default_format,omitempty"`
	// 文本片段列表，可能用于存储组成完整文本的各个部分
	Pieces []*TextPiece `protobuf:"bytes,4,rep,name=pieces,proto3" json:"pieces,omitempty"`
}

func (x *Text) Reset() {
	*x = Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_Text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_data_Text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_data_Text_proto_rawDescGZIP(), []int{0}
}

func (x *Text) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Text) GetDefaultPattern() string {
	if x != nil {
		return x.DefaultPattern
	}
	return ""
}

func (x *Text) GetDefaultFormat() *TextFormat {
	if x != nil {
		return x.DefaultFormat
	}
	return nil
}

func (x *Text) GetPieces() []*TextPiece {
	if x != nil {
		return x.Pieces
	}
	return nil
}

var File_data_Text_proto protoreflect.FileDescriptor

var file_data_Text_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x1a, 0x15,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12,
	0x3e, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x2e, 0x0a, 0x06, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65,
	0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x52, 0x06, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x42,
	0x65, 0x0a, 0x45, 0x63, 0x6f, 0x6f, 0x6c, 0x2e, 0x73, 0x63, 0x78, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x69, 0x6d,
	0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x65, 0x62, 0x63,
	0x61, 0x73, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x01, 0x5a, 0x1a, 0x74, 0x69, 0x6b, 0x74,
	0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_Text_proto_rawDescOnce sync.Once
	file_data_Text_proto_rawDescData = file_data_Text_proto_rawDesc
)

func file_data_Text_proto_rawDescGZIP() []byte {
	file_data_Text_proto_rawDescOnce.Do(func() {
		file_data_Text_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_Text_proto_rawDescData)
	})
	return file_data_Text_proto_rawDescData
}

var file_data_Text_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_Text_proto_goTypes = []interface{}{
	(*Text)(nil),       // 0: tiktok_hack.Text
	(*TextFormat)(nil), // 1: tiktok_hack.TextFormat
	(*TextPiece)(nil),  // 2: tiktok_hack.TextPiece
}
var file_data_Text_proto_depIdxs = []int32{
	1, // 0: tiktok_hack.Text.default_format:type_name -> tiktok_hack.TextFormat
	2, // 1: tiktok_hack.Text.pieces:type_name -> tiktok_hack.TextPiece
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_data_Text_proto_init() }
func file_data_Text_proto_init() {
	if File_data_Text_proto != nil {
		return
	}
	file_data_TextFormat_proto_init()
	file_data_TextPiece_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_Text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text); i {
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
			RawDescriptor: file_data_Text_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_Text_proto_goTypes,
		DependencyIndexes: file_data_Text_proto_depIdxs,
		MessageInfos:      file_data_Text_proto_msgTypes,
	}.Build()
	File_data_Text_proto = out.File
	file_data_Text_proto_rawDesc = nil
	file_data_Text_proto_goTypes = nil
	file_data_Text_proto_depIdxs = nil
}
