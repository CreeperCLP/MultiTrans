package ws

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"clipboard-transfer-station/backend/internal/service"
)


// upgrader 用于将 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}


// Client 表示一个 WebSocket 客户端连接
type Client struct {
	conn *websocket.Conn                 // WebSocket 连接
	send chan service.ClipboardData      // 发送消息通道
}


// Hub 管理所有 WebSocket 客户端和广播
type Hub struct {
	clients    map[*Client]bool              // 所有活跃客户端
	broadcast  chan service.ClipboardData    // 广播消息通道
	register   chan *Client                  // 注册新客户端
	unregister chan *Client                  // 注销客户端
}


// NewHub 创建并启动 WebSocket Hub
func NewHub() *Hub {
       h := &Hub{
	       clients:    make(map[*Client]bool),
	       broadcast:  make(chan service.ClipboardData, 100),
	       register:   make(chan *Client),
	       unregister: make(chan *Client),
       }
       go h.run()
       return h
}


// run 负责管理客户端注册、注销和消息广播
func (h *Hub) run() {
       for {
	       select {
	       case client := <-h.register:
		       h.clients[client] = true
	       case client := <-h.unregister:
		       if _, ok := h.clients[client]; ok {
			       delete(h.clients, client)
			       close(client.send)
		       }
	       case message := <-h.broadcast:
		       for client := range h.clients {
			       select {
			       case client.send <- message:
			       default:
				       close(client.send)
				       delete(h.clients, client)
			       }
		       }
	       }
       }
}


// ServeWs 处理 WebSocket 连接升级和客户端注册
func ServeWs(hub *Hub, clipboardManager *service.ClipboardManager) gin.HandlerFunc {
       return func(c *gin.Context) {
	       conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	       if err != nil {
		       return
	       }
	       client := &Client{conn: conn, send: make(chan service.ClipboardData, 10)}
	       hub.register <- client
	       go writePump(client)
	       // 立即推送当前内容
	       client.send <- clipboardManager.GetContent()
	       // 监听广播并推送给 Hub
	       go func() {
		       for data := range clipboardManager.Broadcast {
			       hub.broadcast <- data
		       }
	       }()
       }
}


// writePump 持续向客户端写入消息，连接关闭时自动清理
func writePump(client *Client) {
       for msg := range client.send {
	       client.conn.WriteJSON(msg)
       }
       client.conn.Close()
}
