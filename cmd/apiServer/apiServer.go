package apiServer

import (
	"strconv"
	"xiliangzi_pro/internal/router"

	"github.com/gin-gonic/gin"
)

func ApiServerInit(port int) {
	// 初始化http router
	r := gin.Default()
	// 初始化路由组
	router.InitRouter(r)
	// 监听端口
	r.Run(":" + strconv.Itoa(port))
}
