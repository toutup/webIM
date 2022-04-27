package websocket

import (
	"net/http"
	"strconv"
	"time"
	"xiliangzi_pro/internal/pkg/middleware/jwt"
	"xiliangzi_pro/internal/websocket/ws"

	"github.com/gin-gonic/gin"
)

func WsServerInit(port int) {

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success", "data": time.Now()})
		return
	})
	go ws.Manager.Start()
	// 加载静态资源
	r.Static("/css", "static/css")
	r.Static("/js", "static/js")
	// 加载模板文件
	r.LoadHTMLGlob("static/html/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
		return
	})
	r.Use(jwt.JWT())
	r.GET("ws", ws.WsServer)
	r.Run(":" + strconv.Itoa(port))
}
