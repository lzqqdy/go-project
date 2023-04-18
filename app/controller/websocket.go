package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ws = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// 允许跨域访问
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handlers 创建业务处理函数映射表
var Handlers = map[string]func(*websocket.Conn){
	"chat": handleChat,
	//  TODO
}

// Websocket Websocket
func Websocket(c *gin.Context) {
	business := c.Param("business")
	conn, err := ws.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	if handler, ok := Handlers[business]; ok {
		handler(conn)
	} else {
		log.Printf("未知业务：%s\n", business)
		err := conn.Close()
		if err != nil {
			return
		}
	}
}

// 处理 chat 业务
func handleChat(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// 处理消息
		log.Printf("chat 收到消息：%s\n", msg)

		// 发送响应
		err = conn.WriteMessage(websocket.TextMessage, []byte("chat 已收到消息"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
