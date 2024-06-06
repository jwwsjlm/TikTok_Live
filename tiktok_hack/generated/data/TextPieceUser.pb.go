// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: data/TextPieceUser.proto

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

type TextPieceUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户信息，可能包含用户的详细信息
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// 是否带冒号，可能表示在显示时用户名后是否需要加上冒号
	WithColon bool `protobuf:"varint,2,opt,name=with_colon,json=withColon,proto3" json:"with_colon,omitempty"`
}

func (x *TextPieceUser) Reset() {
	*x = TextPieceUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_TextPieceUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextPieceUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextPieceUser) ProtoMessage() {}

func (x *TextPieceUser) ProtoReflect() protoreflect.Message {
	mi := &file_data_TextPieceUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextPieceUser.ProtoReflect.Descriptor instead.
func (*TextPieceUser) Descriptor() ([]byte, []int) {
	return file_data_TextPieceUser_proto_rawDescGZIP(), []int{0}
}

func (x *TextPieceUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TextPieceUser) GetWithColon() bool {
	if x != nil {
		return x.WithColon
	}
	return false
}

var File_data_TextPieceUser_proto protoreflect.FileDescriptor

var file_data_TextPieceUser_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x69, 0x6b, 0x74,
	0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x1a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x55, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0d, 0x54, 0x65, 0x78, 0x74,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x77, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6c, 0x6f, 0x6e, 0x42,
	0x1c, 0x5a, 0x1a, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_TextPieceUser_proto_rawDescOnce sync.Once
	file_data_TextPieceUser_proto_rawDescData = file_data_TextPieceUser_proto_rawDesc
)

func file_data_TextPieceUser_proto_rawDescGZIP() []byte {
	file_data_TextPieceUser_proto_rawDescOnce.Do(func() {
		file_data_TextPieceUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_TextPieceUser_proto_rawDescData)
	})
	return file_data_TextPieceUser_proto_rawDescData
}

var file_data_TextPieceUser_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_TextPieceUser_proto_goTypes = []interface{}{
	(*TextPieceUser)(nil), // 0: tiktok_hack.TextPieceUser
	(*User)(nil),          // 1: tiktok_hack.User
}
var file_data_TextPieceUser_proto_depIdxs = []int32{
	1, // 0: tiktok_hack.TextPieceUser.user:type_name -> tiktok_hack.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_data_TextPieceUser_proto_init() }
func file_data_TextPieceUser_proto_init() {
	if File_data_TextPieceUser_proto != nil {
		return
	}
	file_data_User_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_TextPieceUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextPieceUser); i {
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
			RawDescriptor: file_data_TextPieceUser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_TextPieceUser_proto_goTypes,
		DependencyIndexes: file_data_TextPieceUser_proto_depIdxs,
		MessageInfos:      file_data_TextPieceUser_proto_msgTypes,
	}.Build()
	File_data_TextPieceUser_proto = out.File
	file_data_TextPieceUser_proto_rawDesc = nil
	file_data_TextPieceUser_proto_goTypes = nil
	file_data_TextPieceUser_proto_depIdxs = nil
}
