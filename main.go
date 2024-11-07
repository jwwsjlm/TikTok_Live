package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"Sunny/tiktok_hack/generated"
	"github.com/gorilla/websocket"
	"github.com/qtgolang/SunnyNet/SunnyNet"
	"github.com/qtgolang/SunnyNet/public"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var Sunny = SunnyNet.NewSunny()
var agentlist sync.Map // 使用 sync.Map 替代 map[string]*agent

// agent 结构体用于存储 WebSocket 连接及其相关信息
func main() {
	// 绑定回调函数
	Sunny.SetGoCallback(HttpCallback, TcpCallback, WSCallback, UdpCallback)

	// 设置端口并启动
	s := Sunny.SetPort(23809)
	defer s.Close()
	//随机tls指纹
	//s.SetRandomTLS(true)
	s.SetGlobalProxy("socket5://127.0.0.1:21586")
	st := s.Start()
	if st.Error != nil {
		log.Fatalf(st.Error.Error())
	}

	// 设置 WebSocket 升级器
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有 CORS 请求
		},
	}

	// 处理 WebSocket 请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(upgrader, w, r)
	})

	fmt.Println("浏览器代理设置为:127.0.0.1:23809")
	fmt.Println("上游代理地址为:127.0.0.1:21586")
	fmt.Println("正在运行....")

	// 启动 HTTP 服务器
	go func() {
		err := http.ListenAndServe(":18080", nil)
		if err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	log.Println("WebSocket 服务启动成功.ws端口为18080")

	// 避免程序退出
	select {}
}

// handleWebSocket 处理 WebSocket 连接
func handleWebSocket(upgrader websocket.Upgrader, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 升级失败:", err)
		return
	}
	defer conn.Close()

	key := r.Header.Get("Sec-WebSocket-Key")
	if key == "" {
		log.Println("缺少 Sec-WebSocket-Key 头")
		return
	}

	agent := &agent{
		queue: make(chan []byte, 2),
		stop:  make(chan struct{}),
		conn:  conn,
	}
	agentlist.Store(key, agent)
	go handleClient(agent)

	log.Println("当前连接数", getAgentCount())

	// 设置读取超时时间
	err = agent.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		log.Println("设置读取超时失败:", err)
		return
	}

	defer func() {
		log.Println(key, "断开连接")
		close(agent.stop)
		agentlist.Delete(key)
	}()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息失败:", err)
			break
		}
		log.Printf("收到消息: %s", message)
		if string(message) == "ping" {
			if err := conn.WriteMessage(mt, []byte("pong")); err != nil {
				log.Println("发送消息失败:", err)
				break
			} else {
				_ = agent.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
			}
		}
	}
}

// handleClient 处理客户端消息
func handleClient(a *agent) {
	for {
		select {
		case <-a.stop:
			return
		case b := <-a.queue:
			err := a.conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("发送消息失败:", err)
			}
		}
	}
}

// getAgentCount 获取当前连接数
func getAgentCount() int {
	count := 0
	agentlist.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}

// HttpCallback HTTP 回调函数
func HttpCallback(Conn *SunnyNet.HttpConn) {
	// 处理 HTTP 连接
}

// WSCallback WebSocket 回调函数
func WSCallback(Conn *SunnyNet.WsConn) {
	if !strings.Contains(Conn.Url, "tiktok.com/webcast/im/") {
		return
	}

	message := Conn.GetMessageBody()
	PushFrame := &tiktok_hack.WebcastPushFrame{}
	err := proto.Unmarshal(message, PushFrame)
	if err != nil {
		log.Println("解析消息失败:", err)
		return
	}

	if PushFrame.PayloadType == "ack" {
		return // 心跳包数据不处理
	}

	isGzip := CheckGzip(PushFrame)
	if isGzip && PushFrame.PayloadType == "msg" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(PushFrame.Payload))
		if err != nil {
			log.Println("解析 Gzip 消息失败:", err)
			return
		}
		defer gzipReader.Close()

		uncompressedData, err := io.ReadAll(gzipReader)
		if err != nil {
			log.Println("读取解压数据失败:", err)
			return
		}

		response := &tiktok_hack.WebcastResponse{}
		err = proto.Unmarshal(uncompressedData, response)
		if err != nil {
			log.Println("解析解压数据失败:", err)
			return
		}

		for _, v := range response.Messages {
			msg, err := MatchMethod(v.Method)
			if err != nil {
				log.Printf("未知消息，无法处理: %v, %s\n", err, hex.EncodeToString(v.Payload))
				continue
			}
			err = proto.Unmarshal(v.Payload, msg)
			if err != nil {
				log.Println("解析消息失败:", err)
				continue
			}

			// 序列化为 JSON
			marshal, err := protojson.Marshal(msg)
			if err != nil {
				log.Println("JSON 序列化失败:", err)
				continue
			}

			// 遍历 agentlist，将消息发送到每个连接
			agentlist.Range(func(_, value interface{}) bool {
				agent := value.(*agent)
				agent.queue <- marshal
				return true
			})
		}

	}
}

// CheckGzip 检查协议头当中是否包含gzip
func CheckGzip(headers *tiktok_hack.WebcastPushFrame) bool {
	return headers.Headers["compress_type"] == "gzip"
}

// TcpCallback TCP 回调函数
func TcpCallback(Conn *SunnyNet.TcpConn) {
	// 处理 TCP 连接
}

// UdpCallback UDP 回调函数
func UdpCallback(Conn *SunnyNet.UDPConn) {
	if public.SunnyNetUDPTypeReceive == Conn.Type {
		// 处理接收的 UDP 数据
	}
	if public.SunnyNetUDPTypeSend == Conn.Type {
		// 处理发送的 UDP 数据
	}
	if public.SunnyNetUDPTypeClosed == Conn.Type {
		// 处理关闭的 UDP 连接
	}
}
func MatchMethod(method string) (protoreflect.ProtoMessage, error) {
	if createMessage, ok := messageTypeMap[method]; ok {
		return createMessage(), nil
	}
	return nil, errors.New("未知消息: " + method)
}
