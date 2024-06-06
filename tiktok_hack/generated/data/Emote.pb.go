// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: data/Emote.proto

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

type Emote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 表情ID，唯一标识一个表情
	EmoteId string `protobuf:"bytes,1,opt,name=emote_id,json=emoteId,proto3" json:"emote_id,omitempty"`
	// 表情图像，可能包含表情的图像信息
	Image *Image `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// 审核状态，可能表示表情的审核结果或状态
	AuditStatus int32 `protobuf:"varint,3,opt,name=audit_status,json=auditStatus,proto3" json:"audit_status,omitempty"`
	// 唯一标识符，可能用于唯一标识表情的UUID
	Uuid string `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// 表情类型，可能用于分类表情
	EmoteType int32 `protobuf:"varint,5,opt,name=emote_type,json=emoteType,proto3" json:"emote_type,omitempty"`
	// 内容来源，可能表示表情内容的来源
	ContentSource int32 `protobuf:"varint,6,opt,name=content_source,json=contentSource,proto3" json:"content_source,omitempty"`
	// 表情私有类型，可能用于表示表情的私有分类或属性
	EmotePrivateType int32 `protobuf:"varint,7,opt,name=emote_private_type,json=emotePrivateType,proto3" json:"emote_private_type,omitempty"`
	// 包ID，可能用于标识表情所属的包或集合
	PackageId string `protobuf:"bytes,8,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	// 审核信息，可能包含更详细的审核结果或相关信息
	AuditInfo *AuditInfo `protobuf:"bytes,9,opt,name=audit_info,json=auditInfo,proto3" json:"audit_info,omitempty"`
	// 奖励条件，可能表示与表情相关的奖励或解锁条件
	RewardCondition int32 `protobuf:"varint,10,opt,name=reward_condition,json=rewardCondition,proto3" json:"reward_condition,omitempty"`
	// 表情上传信息，可能包含表情上传的详细数据
	EmoteUploadInfo *EmoteUploadInfo `protobuf:"bytes,11,opt,name=emote_upload_info,json=emoteUploadInfo,proto3" json:"emote_upload_info,omitempty"`
	// 创建时间，可能表示表情创建的时间
	CreateTime int64 `protobuf:"varint,12,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Emote) Reset() {
	*x = Emote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_Emote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emote) ProtoMessage() {}

func (x *Emote) ProtoReflect() protoreflect.Message {
	mi := &file_data_Emote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emote.ProtoReflect.Descriptor instead.
func (*Emote) Descriptor() ([]byte, []int) {
	return file_data_Emote_proto_rawDescGZIP(), []int{0}
}

func (x *Emote) GetEmoteId() string {
	if x != nil {
		return x.EmoteId
	}
	return ""
}

func (x *Emote) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Emote) GetAuditStatus() int32 {
	if x != nil {
		return x.AuditStatus
	}
	return 0
}

func (x *Emote) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Emote) GetEmoteType() int32 {
	if x != nil {
		return x.EmoteType
	}
	return 0
}

func (x *Emote) GetContentSource() int32 {
	if x != nil {
		return x.ContentSource
	}
	return 0
}

func (x *Emote) GetEmotePrivateType() int32 {
	if x != nil {
		return x.EmotePrivateType
	}
	return 0
}

func (x *Emote) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *Emote) GetAuditInfo() *AuditInfo {
	if x != nil {
		return x.AuditInfo
	}
	return nil
}

func (x *Emote) GetRewardCondition() int32 {
	if x != nil {
		return x.RewardCondition
	}
	return 0
}

func (x *Emote) GetEmoteUploadInfo() *EmoteUploadInfo {
	if x != nil {
		return x.EmoteUploadInfo
	}
	return nil
}

func (x *Emote) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_data_Emote_proto protoreflect.FileDescriptor

var file_data_Emote_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x1a,
	0x10, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x45, 0x6d,
	0x6f, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03, 0x0a, 0x05, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0a,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48,
	0x0a, 0x11, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x69, 0x6b, 0x74,
	0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x74, 0x69, 0x6b,
	0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_Emote_proto_rawDescOnce sync.Once
	file_data_Emote_proto_rawDescData = file_data_Emote_proto_rawDesc
)

func file_data_Emote_proto_rawDescGZIP() []byte {
	file_data_Emote_proto_rawDescOnce.Do(func() {
		file_data_Emote_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_Emote_proto_rawDescData)
	})
	return file_data_Emote_proto_rawDescData
}

var file_data_Emote_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_Emote_proto_goTypes = []interface{}{
	(*Emote)(nil),           // 0: tiktok_hack.Emote
	(*Image)(nil),           // 1: tiktok_hack.Image
	(*AuditInfo)(nil),       // 2: tiktok_hack.AuditInfo
	(*EmoteUploadInfo)(nil), // 3: tiktok_hack.EmoteUploadInfo
}
var file_data_Emote_proto_depIdxs = []int32{
	1, // 0: tiktok_hack.Emote.image:type_name -> tiktok_hack.Image
	2, // 1: tiktok_hack.Emote.audit_info:type_name -> tiktok_hack.AuditInfo
	3, // 2: tiktok_hack.Emote.emote_upload_info:type_name -> tiktok_hack.EmoteUploadInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_data_Emote_proto_init() }
func file_data_Emote_proto_init() {
	if File_data_Emote_proto != nil {
		return
	}
	file_data_Image_proto_init()
	file_data_AuditInfo_proto_init()
	file_data_EmoteUploadInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_Emote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emote); i {
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
			RawDescriptor: file_data_Emote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_Emote_proto_goTypes,
		DependencyIndexes: file_data_Emote_proto_depIdxs,
		MessageInfos:      file_data_Emote_proto_msgTypes,
	}.Build()
	File_data_Emote_proto = out.File
	file_data_Emote_proto_rawDesc = nil
	file_data_Emote_proto_goTypes = nil
	file_data_Emote_proto_depIdxs = nil
}
