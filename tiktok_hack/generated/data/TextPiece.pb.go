// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: data/TextPiece.proto

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
type TextPiece struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 类型，可能表示文本片段的类型，例如普通文本、链接、特殊标记等
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// 格式，包含文本片段的格式信息，如颜色、大小等
	Format *TextFormat `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	// 字符串值，对于某些类型的文本片段，这里可能存储实际的字符串文本
	StringValue string `protobuf:"bytes,11,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	// 用户值，可能包含与用户相关的信息，如用户名或用户ID
	UserValue *TextPieceUser `protobuf:"bytes,21,opt,name=user_value,json=userValue,proto3" json:"user_value,omitempty"`
	// 礼物值，可能包含礼物信息，如礼物名称或ID
	GiftValue *TextPieceGift `protobuf:"bytes,22,opt,name=gift_value,json=giftValue,proto3" json:"gift_value,omitempty"`
	// 爱心值，可能包含与爱心相关的信息，如爱心的数量或类型
	HeartValue *TextPieceHeart `protobuf:"bytes,23,opt,name=heart_value,json=heartValue,proto3" json:"heart_value,omitempty"`
	// 模式引用值，可能包含对预定义模式的引用
	PatternRefValue *TextPiecePatternRef `protobuf:"bytes,24,opt,name=pattern_ref_value,json=patternRefValue,proto3" json:"pattern_ref_value,omitempty"`
	// 图像值，可能包含与图像相关的信息，如图像URL或图像ID
	ImageValue *TextPieceImage `protobuf:"bytes,25,opt,name=image_value,json=imageValue,proto3" json:"image_value,omitempty"`
}

func (x *TextPiece) Reset() {
	*x = TextPiece{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_TextPiece_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextPiece) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextPiece) ProtoMessage() {}

func (x *TextPiece) ProtoReflect() protoreflect.Message {
	mi := &file_data_TextPiece_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextPiece.ProtoReflect.Descriptor instead.
func (*TextPiece) Descriptor() ([]byte, []int) {
	return file_data_TextPiece_proto_rawDescGZIP(), []int{0}
}

func (x *TextPiece) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TextPiece) GetFormat() *TextFormat {
	if x != nil {
		return x.Format
	}
	return nil
}

func (x *TextPiece) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

func (x *TextPiece) GetUserValue() *TextPieceUser {
	if x != nil {
		return x.UserValue
	}
	return nil
}

func (x *TextPiece) GetGiftValue() *TextPieceGift {
	if x != nil {
		return x.GiftValue
	}
	return nil
}

func (x *TextPiece) GetHeartValue() *TextPieceHeart {
	if x != nil {
		return x.HeartValue
	}
	return nil
}

func (x *TextPiece) GetPatternRefValue() *TextPiecePatternRef {
	if x != nil {
		return x.PatternRefValue
	}
	return nil
}

func (x *TextPiece) GetImageValue() *TextPieceImage {
	if x != nil {
		return x.ImageValue
	}
	return nil
}

var File_data_TextPiece_proto protoreflect.FileDescriptor

var file_data_TextPiece_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68,
	0x61, 0x63, 0x6b, 0x1a, 0x15, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x50,
	0x69, 0x65, 0x63, 0x65, 0x47, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x52, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x03, 0x0a, 0x09, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f,
	0x68, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x78,
	0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x69, 0x6b, 0x74,
	0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63,
	0x65, 0x47, 0x69, 0x66, 0x74, 0x52, 0x09, 0x67, 0x69, 0x66, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x3c, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68,
	0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4c,
	0x0a, 0x11, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x69, 0x6b, 0x74,
	0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63,
	0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x65, 0x66, 0x52, 0x0f, 0x70, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x52, 0x65, 0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e,
	0x54, 0x65, 0x78, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0a,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x65, 0x0a, 0x45, 0x63, 0x6f,
	0x6f, 0x6c, 0x2e, 0x73, 0x63, 0x78, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x74, 0x69,
	0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x65, 0x62, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x50, 0x01, 0x5a, 0x1a, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61,
	0x63, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_TextPiece_proto_rawDescOnce sync.Once
	file_data_TextPiece_proto_rawDescData = file_data_TextPiece_proto_rawDesc
)

func file_data_TextPiece_proto_rawDescGZIP() []byte {
	file_data_TextPiece_proto_rawDescOnce.Do(func() {
		file_data_TextPiece_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_TextPiece_proto_rawDescData)
	})
	return file_data_TextPiece_proto_rawDescData
}

var file_data_TextPiece_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_TextPiece_proto_goTypes = []interface{}{
	(*TextPiece)(nil),           // 0: tiktok_hack.TextPiece
	(*TextFormat)(nil),          // 1: tiktok_hack.TextFormat
	(*TextPieceUser)(nil),       // 2: tiktok_hack.TextPieceUser
	(*TextPieceGift)(nil),       // 3: tiktok_hack.TextPieceGift
	(*TextPieceHeart)(nil),      // 4: tiktok_hack.TextPieceHeart
	(*TextPiecePatternRef)(nil), // 5: tiktok_hack.TextPiecePatternRef
	(*TextPieceImage)(nil),      // 6: tiktok_hack.TextPieceImage
}
var file_data_TextPiece_proto_depIdxs = []int32{
	1, // 0: tiktok_hack.TextPiece.format:type_name -> tiktok_hack.TextFormat
	2, // 1: tiktok_hack.TextPiece.user_value:type_name -> tiktok_hack.TextPieceUser
	3, // 2: tiktok_hack.TextPiece.gift_value:type_name -> tiktok_hack.TextPieceGift
	4, // 3: tiktok_hack.TextPiece.heart_value:type_name -> tiktok_hack.TextPieceHeart
	5, // 4: tiktok_hack.TextPiece.pattern_ref_value:type_name -> tiktok_hack.TextPiecePatternRef
	6, // 5: tiktok_hack.TextPiece.image_value:type_name -> tiktok_hack.TextPieceImage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_data_TextPiece_proto_init() }
func file_data_TextPiece_proto_init() {
	if File_data_TextPiece_proto != nil {
		return
	}
	file_data_TextFormat_proto_init()
	file_data_TextPieceUser_proto_init()
	file_data_TextPieceGift_proto_init()
	file_data_TextPieceHeart_proto_init()
	file_data_TextPiecePatternRef_proto_init()
	file_data_TextPieceImage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_TextPiece_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextPiece); i {
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
			RawDescriptor: file_data_TextPiece_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_TextPiece_proto_goTypes,
		DependencyIndexes: file_data_TextPiece_proto_depIdxs,
		MessageInfos:      file_data_TextPiece_proto_msgTypes,
	}.Build()
	File_data_TextPiece_proto = out.File
	file_data_TextPiece_proto_rawDesc = nil
	file_data_TextPiece_proto_goTypes = nil
	file_data_TextPiece_proto_depIdxs = nil
}
