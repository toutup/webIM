package ws

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"xiliangzi_pro/pkg/online"
	"xiliangzi_pro/pkg/util"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wu = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 初始化websocket连接
func WsServer(c *gin.Context) {
	// 获取用户连接信息
	user := util.GetUserInfo(c)
	uid, _ := strconv.Atoi(user["uid"])
	wu.Subprotocols = []string{c.Request.Header.Get("Sec-WebSocket-Protocol")}
	// 创建连接
	conn, err := wu.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("websocket connect error:%s", c.Param("channel")))
		http.NotFound(c.Writer, c.Request)
		return
	}
	// 生成唯一标识client_id
	var cleintId = uuid.NewV4().String()
	client := &Client{
		Id:     cleintId,
		Socket: conn,
		Send:   make(chan []byte),
	}
	// 将注册信息存储到channel
	Manager.Online <- client
	// 新加入用户加入在线列表
	onlineServer := online.NewOnline()
	req := online.AddOnlineListRequest{
		Prefix:   "online",
		SetName:  "list",
		Uid:      uid,
		DateTime: time.Now().Unix(),
		Data:     user,
	}
	err = onlineServer.AddOnlineList(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 监听客户端和服务端消息
	go client.Read()
	go client.Write()
}
