package groups

import (
	v1 "xiliangzi_pro/api/auth/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	// 根据用户id获取用户信息
	r.GET("user/get_user_by_id/:id", v1.GetUserById)
	r.GET("user/ws", v1.GetUserById)
}
