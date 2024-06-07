package main

import (
	tiktokhack "Sunny/tiktok_hack/generated"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/qtgolang/SunnyNet/SunnyNet"
	"github.com/qtgolang/SunnyNet/public"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
	"log"
	"strings"
)

var Sunny = SunnyNet.NewSunny()

func main() {
	//绑定回调函数
	Sunny.SetGoCallback(HttpCallback, TcpCallback, WSCallback, UdpCallback)

	fmt.Println("正在运行....")
	//绑定端口号并启动

	s := Sunny.SetPort(10808)
	//随机tls指纹
	//s.SetRandomTLS(true)
	//设置全局上游代理 仅支持Socket5和http 例如 socket5://admin:123456@127.0.0.1:8888 或 http://admin:123456@127.0.0.1:8888
	s.SetGlobalProxy("socket5://127.0.0.1:30801") //这个是全局的
	s.Start()

	//避免程序退出
	select {}

}

func HttpCallback(Conn *SunnyNet.HttpConn) {

}

func WSCallback(Conn *SunnyNet.WsConn) {
	//tiktok.com/webcast/im/
	if strings.Contains(Conn.Url, "webcast/im/") == false {
		return

	}
	//捕获到数据可以修改,修改空数据,取消发送/接收
	message := Conn.GetMessageBody()
	//fmt.Println("ws数据", Conn.GetMessageBody())
	PushFrame := &tiktokhack.WebcastPushFrame{}
	err := proto.Unmarshal(message, PushFrame)
	if err != nil {
		log.Println("解析消息失败：", err)
		return
	}
	//log.Println(PushFrame.Payload)
	if PushFrame.PayloadType == "ack" {
		//心跳包数据不做处理
		return
	}
	n := false
	for v, p := range PushFrame.Headers {
		if v == "compress_type" {
			if p == "gzip" {
				n = true
				continue
			}
		}
	}
	//消息为gzip压缩
	if n == true && PushFrame.PayloadType == "msg" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(PushFrame.Payload))
		defer gzipReader.Close()
		if err != nil {
			log.Println("解析消息失败gzip：", err, PushFrame.PayloadType)
			return
		}

		uncompressedData, _ := io.ReadAll(gzipReader)
		response := &tiktokhack.WebcastResponse{}
		err = proto.Unmarshal(uncompressedData, response)

		for _, v := range response.Messages {
			msg, err := Match(v.Method)
			if err != nil {
				log.Println("未知的消息类型", err, PushFrame.PayloadType)
				continue
			}
			err = proto.Unmarshal(v.Payload, msg)
			if err != nil {
				log.Println("解析消息失败3：", err)
				continue
			}
			marshal, err := protojson.Marshal(msg)
			if err != nil {
				log.Println("protojson:unmarshal:", err)
				return
			}
			log.Println(string(marshal))

		}
		//log.Println(response.)
	}

}
func Match(Method string) (protoreflect.ProtoMessage, error) {
	switch Method {
	case "WebcastChatMessage":
		return &tiktokhack.WebcastChatMessage{}, nil
	case "WebcastMemberMessage":
		return &tiktokhack.WebcastMemberMessage{}, nil
	case "WebcastRoomUserSeqMessage":
		return &tiktokhack.WebcastRoomUserSeqMessage{}, nil
	case "WebcastLikeMessage":
		return &tiktokhack.WebcastLikeMessage{}, nil
	case "WebcastSocialMessage":
		return &tiktokhack.WebcastSocialMessage{}, nil
	case "WebcastGiftMessage":
		return &tiktokhack.WebcastGiftMessage{}, nil
	case "WebcastImDeleteMessage":
		return &tiktokhack.WebcastImDeleteMessage{}, nil
	case "WebcastUnauthorizedMemberMessage":
		return &tiktokhack.WebcastUnauthorizedMemberMessage{}, nil
	case "WebcastRankUpdateMessage":
		return &tiktokhack.WebcastRankUpdateMessage{}, nil
	case "WebcastLinkMicArmies":
		return &tiktokhack.WebcastLinkMicArmies{}, nil

	default:
		return nil, fmt.Errorf("未知的消息类型:" + Method)
	}
}
func TcpCallback(Conn *SunnyNet.TcpConn) {
	//捕获到数据可以修改,修改空数据,取消发送/接收
	//Conn.SetAgent("") //这个是针对这个一个请求的,TCP,只能设置S5代理,
	//fmt.Println(Conn.Pid, Conn.LocalAddr, Conn.RemoteAddr, Conn.Type, Conn.GetBodyLen())
}
func UdpCallback(Conn *SunnyNet.UDPConn) {
	//在 Windows 捕获UDP需要加载驱动,并且设置进程名
	//其他情况需要设置Socket5代理,才能捕获到UDP
	//捕获到数据可以修改,修改空数据,取消发送/接收
	if public.SunnyNetUDPTypeReceive == Conn.Type {
		//fmt.Println("接收UDP", Conn.LocalAddress, Conn.RemoteAddress, len(Conn.Data))
	}
	if public.SunnyNetUDPTypeSend == Conn.Type {
		//fmt.Println("发送UDP", Conn.LocalAddress, Conn.RemoteAddress, len(Conn.Data))
	}
	if public.SunnyNetUDPTypeClosed == Conn.Type {
		//fmt.Println("关闭UDP", Conn.LocalAddress, Conn.RemoteAddress)
	}
}
