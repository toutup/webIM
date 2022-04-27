package groups

import (
	v1 "xiliangzi_pro/api/auth/v1"

	"github.com/gin-gonic/gin"
)

// 初始化auth相关路由
func AuthRouterInit(router *gin.RouterGroup) {

	api := router.Group("auth/")
	api.POST("register", v1.Register)
	api.POST("login", v1.Login)
}
