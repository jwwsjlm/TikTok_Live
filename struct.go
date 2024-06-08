package main

import (
	"github.com/gorilla/websocket"
)

type agent struct {
	queue chan []byte     // 消息队列
	conn  *websocket.Conn // websocket连接
	stop  chan struct{}
}
