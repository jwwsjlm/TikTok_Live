package main

import (
	tiktokhack "Sunny/tiktok_hack/generated"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type agent struct {
	queue chan []byte     // 消息队列
	conn  *websocket.Conn // websocket连接
	stop  chan struct{}
}

var messageTypeMap = map[string]func() protoreflect.ProtoMessage{
	"WebcastChatMessage":               func() protoreflect.ProtoMessage { return &tiktokhack.WebcastChatMessage{} },
	"WebcastMemberMessage":             func() protoreflect.ProtoMessage { return &tiktokhack.WebcastMemberMessage{} },
	"WebcastRoomUserSeqMessage":        func() protoreflect.ProtoMessage { return &tiktokhack.WebcastRoomUserSeqMessage{} },
	"WebcastLikeMessage":               func() protoreflect.ProtoMessage { return &tiktokhack.WebcastLikeMessage{} },
	"WebcastSocialMessage":             func() protoreflect.ProtoMessage { return &tiktokhack.WebcastSocialMessage{} },
	"WebcastGiftMessage":               func() protoreflect.ProtoMessage { return &tiktokhack.WebcastGiftMessage{} },
	"WebcastImDeleteMessage":           func() protoreflect.ProtoMessage { return &tiktokhack.WebcastImDeleteMessage{} },
	"WebcastUnauthorizedMemberMessage": func() protoreflect.ProtoMessage { return &tiktokhack.WebcastUnauthorizedMemberMessage{} },
	"WebcastRankUpdateMessage":         func() protoreflect.ProtoMessage { return &tiktokhack.WebcastRankUpdateMessage{} },
	"WebcastLinkMicArmies":             func() protoreflect.ProtoMessage { return &tiktokhack.WebcastLinkMicArmies{} },
}
