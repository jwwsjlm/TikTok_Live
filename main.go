package main

import (
	Tk "Sunny/tiktok_hack/generated/im"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/qtgolang/SunnyNet/SunnyNet"
	"github.com/qtgolang/SunnyNet/public"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
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
	//捕获到数据可以修改,修改空数据,取消发送/接收
	message := Conn.GetMessageBody()
	//fmt.Println("ws数据", Conn.GetMessageBody())
	PushFrame := &Tk.PushFrame{}
	err := proto.Unmarshal(message, PushFrame)
	if err != nil {
		log.Println("解析消息失败：", err)
		return
	}
	//log.Println(PushFrame.Payload)
	gzipReader, err := gzip.NewReader(bytes.NewReader(PushFrame.Payload))
	if err != nil {
		log.Println("解析消息失败gzip：", err)
		return

	}
	uncompressedData, _ := io.ReadAll(gzipReader)
	response := &Tk.Response{}
	err = proto.Unmarshal(uncompressedData, response)
	for _, v := range response.MessagesList {
		//log.Println(v.Method)
		if v.Method == "WebcastChatMessage" {
			msg := &Tk.ChatMessage{}
			err := proto.Unmarshal(v.Payload, msg)
			if err != nil {
				log.Println("解析消息失败3：", err)
				continue
			}
			log.Printf("用户名:%s,消息内容:%s\n\n", msg.User.Nickname, msg.Content)
		}

	}
	//log.Println(response.)

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
