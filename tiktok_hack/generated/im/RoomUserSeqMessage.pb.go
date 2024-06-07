// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: im/RoomUserSeqMessage.proto

package im

import (
	data "Sunny/tiktok_hack/generated/data"
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

// 更新时间 : 2024-06-6
type RoomUserSeqMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 假设Common是另一个文件中定义的message
	Common     *Common        `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	Ranks      []*Contributor `protobuf:"bytes,2,rep,name=ranks,proto3" json:"ranks,omitempty"`
	Total      string         `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"` // 假设total字段是字符串类型，用于存储int64的文本表示
	PopStr     string         `protobuf:"bytes,4,opt,name=pop_str,json=popStr,proto3" json:"pop_str,omitempty"`
	Seats      []*Contributor `protobuf:"bytes,5,rep,name=seats,proto3" json:"seats,omitempty"`
	Popularity string         `protobuf:"bytes,6,opt,name=popularity,proto3" json:"popularity,omitempty"`                // 假设popularity字段是字符串类型，用于存储int64的文本表示
	TotalUser  string         `protobuf:"bytes,7,opt,name=total_user,json=totalUser,proto3" json:"total_user,omitempty"` // 假设total_user字段是字符串类型，用于存储int64的文本表示
	Anonymous  string         `protobuf:"bytes,8,opt,name=anonymous,proto3" json:"anonymous,omitempty"`                  // 假设anonymous字段是字符串类型，用于存储int64的文本表示
}

func (x *RoomUserSeqMessage) Reset() {
	*x = RoomUserSeqMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_RoomUserSeqMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomUserSeqMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomUserSeqMessage) ProtoMessage() {}

func (x *RoomUserSeqMessage) ProtoReflect() protoreflect.Message {
	mi := &file_im_RoomUserSeqMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomUserSeqMessage.ProtoReflect.Descriptor instead.
func (*RoomUserSeqMessage) Descriptor() ([]byte, []int) {
	return file_im_RoomUserSeqMessage_proto_rawDescGZIP(), []int{0}
}

func (x *RoomUserSeqMessage) GetCommon() *Common {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *RoomUserSeqMessage) GetRanks() []*Contributor {
	if x != nil {
		return x.Ranks
	}
	return nil
}

func (x *RoomUserSeqMessage) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *RoomUserSeqMessage) GetPopStr() string {
	if x != nil {
		return x.PopStr
	}
	return ""
}

func (x *RoomUserSeqMessage) GetSeats() []*Contributor {
	if x != nil {
		return x.Seats
	}
	return nil
}

func (x *RoomUserSeqMessage) GetPopularity() string {
	if x != nil {
		return x.Popularity
	}
	return ""
}

func (x *RoomUserSeqMessage) GetTotalUser() string {
	if x != nil {
		return x.TotalUser
	}
	return ""
}

func (x *RoomUserSeqMessage) GetAnonymous() string {
	if x != nil {
		return x.Anonymous
	}
	return ""
}

type Contributor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score int64      `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"` // 假设score字段是64位整数
	User  *data.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`    // User消息，需要根据User的定义调整字段类型
	Rank  int64      `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`   // 假设rank字段是64位整数
	Delta int64      `protobuf:"varint,4,opt,name=delta,proto3" json:"delta,omitempty"` // 假设delta字段是64位整数
}

func (x *Contributor) Reset() {
	*x = Contributor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_RoomUserSeqMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contributor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contributor) ProtoMessage() {}

func (x *Contributor) ProtoReflect() protoreflect.Message {
	mi := &file_im_RoomUserSeqMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contributor.ProtoReflect.Descriptor instead.
func (*Contributor) Descriptor() ([]byte, []int) {
	return file_im_RoomUserSeqMessage_proto_rawDescGZIP(), []int{1}
}

func (x *Contributor) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Contributor) GetUser() *data.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Contributor) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Contributor) GetDelta() int64 {
	if x != nil {
		return x.Delta
	}
	return 0
}

var File_im_RoomUserSeqMessage_proto protoreflect.FileDescriptor

var file_im_RoomUserSeqMessage_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6d, 0x2f, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x71,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74,
	0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x1a, 0x0f, 0x69, 0x6d, 0x2f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a,
	0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63,
	0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x70, 0x5f, 0x73, 0x74,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x70, 0x53, 0x74, 0x72, 0x12,
	0x2e, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x22, 0x74, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x65, 0x6c,
	0x74, 0x61, 0x42, 0x67, 0x0a, 0x43, 0x63, 0x6f, 0x6f, 0x6c, 0x2e, 0x73, 0x63, 0x78, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x77,
	0x65, 0x62, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x69, 0x6d, 0x50, 0x01, 0x5a, 0x1e, 0x53, 0x75, 0x6e,
	0x6e, 0x79, 0x2f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x68, 0x61, 0x63, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_im_RoomUserSeqMessage_proto_rawDescOnce sync.Once
	file_im_RoomUserSeqMessage_proto_rawDescData = file_im_RoomUserSeqMessage_proto_rawDesc
)

func file_im_RoomUserSeqMessage_proto_rawDescGZIP() []byte {
	file_im_RoomUserSeqMessage_proto_rawDescOnce.Do(func() {
		file_im_RoomUserSeqMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_RoomUserSeqMessage_proto_rawDescData)
	})
	return file_im_RoomUserSeqMessage_proto_rawDescData
}

var file_im_RoomUserSeqMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_im_RoomUserSeqMessage_proto_goTypes = []interface{}{
	(*RoomUserSeqMessage)(nil), // 0: tiktok_hack.RoomUserSeqMessage
	(*Contributor)(nil),        // 1: tiktok_hack.Contributor
	(*Common)(nil),             // 2: tiktok_hack.Common
	(*data.User)(nil),          // 3: tiktok_hack.User
}
var file_im_RoomUserSeqMessage_proto_depIdxs = []int32{
	2, // 0: tiktok_hack.RoomUserSeqMessage.common:type_name -> tiktok_hack.Common
	1, // 1: tiktok_hack.RoomUserSeqMessage.ranks:type_name -> tiktok_hack.Contributor
	1, // 2: tiktok_hack.RoomUserSeqMessage.seats:type_name -> tiktok_hack.Contributor
	3, // 3: tiktok_hack.Contributor.user:type_name -> tiktok_hack.User
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_im_RoomUserSeqMessage_proto_init() }
func file_im_RoomUserSeqMessage_proto_init() {
	if File_im_RoomUserSeqMessage_proto != nil {
		return
	}
	file_im_Common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_im_RoomUserSeqMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomUserSeqMessage); i {
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
		file_im_RoomUserSeqMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contributor); i {
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
			RawDescriptor: file_im_RoomUserSeqMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_im_RoomUserSeqMessage_proto_goTypes,
		DependencyIndexes: file_im_RoomUserSeqMessage_proto_depIdxs,
		MessageInfos:      file_im_RoomUserSeqMessage_proto_msgTypes,
	}.Build()
	File_im_RoomUserSeqMessage_proto = out.File
	file_im_RoomUserSeqMessage_proto_rawDesc = nil
	file_im_RoomUserSeqMessage_proto_goTypes = nil
	file_im_RoomUserSeqMessage_proto_depIdxs = nil
}
