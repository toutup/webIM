package ws

import (
	"encoding/json"
	"log"
	"xiliangzi_pro/pkg/online"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

var Online *online.Online

// 客户端
type Client struct {
	Id     string
	Uid    int
	Socket *websocket.Conn
	Send   chan []byte
}

// 客户端管理
type ClientManager struct {
	Online    chan *Client
	Offline   chan *Client
	BroadCast chan []byte
	Clients   map[string]*Client
}

// 消息定义
type Message struct {
	MsgId     string `json:"msg_id"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

var Manager = ClientManager{
	Online:    make(chan *Client),
	Offline:   make(chan *Client),
	Clients:   make(map[string]*Client),
	BroadCast: make(chan []byte),
}

// 读取客户端消息
func (c *Client) Read() {
	defer func() {
		Manager.Offline <- c
		c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			log.Printf("读取客户端信息错误：%s", err.Error())
			Manager.Offline <- c
			c.Socket.Close()
			break
		}
		log.Printf("读取到客户端信息：%s", string(message))
		Manager.BroadCast <- message
	}
}

// 向客户端发送消息
func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				log.Println("not message")
				c.Socket.WriteMessage(websocket.TextMessage, []byte{})
				return
			}
			log.Printf("发送到客户端的消息：%s", string(msg))
			c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

// 服务端轮询消息
func (m *ClientManager) Start() {
	for {
		log.Printf("管道通信")
		select {
		case conn := <-Manager.Online:
			log.Printf("新用户加入：%s", conn.Id)
			Manager.Clients[conn.Id] = conn
			jsonMessage, _ := json.Marshal(&Message{MsgId: uuid.NewV4().String(), Sender: "系统消息", Recipient: conn.Id, Content: "Successful connection to socket service"})
			conn.Send <- jsonMessage
		case conn := <-Manager.Offline:
			log.Printf("用户离开：%s", conn.Id)
			if _, ok := Manager.Clients[conn.Id]; ok {
				jsonMessage, _ := json.Marshal(&Message{MsgId: uuid.NewV4().String(), Sender: "系统消息", Recipient: conn.Id, Content: "A socket has disconnected"})
				conn.Send <- jsonMessage
				close(conn.Send)
				delete(Manager.Clients, conn.Id)
			}
		case message := <-Manager.BroadCast:
			log.Printf("发送消息：%s", string(message))
			MsgStruct := Message{}
			json.Unmarshal(message, &MsgStruct)
			for _, conn := range Manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.Id)
				}
			}
		}
	}
}
