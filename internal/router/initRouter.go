package router

import (
	"fmt"
	"net/http"
	"time"
	"xiliangzi_pro/internal/pkg/middleware/jwt"
	groups2 "xiliangzi_pro/internal/router/groups"

	"github.com/gin-gonic/gin"
)

// 初始化各模块路由
func InitRouter(r *gin.Engine) {
	// 绑定路由规则，执行函数
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "health - "+fmt.Sprint(time.Now().Unix()))
	})

	// 加载auth路由
	apiv1 := r.Group("/api/v1/")
	groups2.AuthRouterInit(apiv1)
	// token鉴权中间件
	apiv1.Use(jwt.JWT())
	// 初始化用户路由
	groups2.InitUserRouter(apiv1)
}
