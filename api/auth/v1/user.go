package v1

import (
	"net/http"
	"xiliangzi_pro/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 根据用户id获取用户信息
func GetUserById(c *gin.Context) {
	// 接收参数
	uid := com.StrTo(c.Param("id")).MustInt()
	// 初始化userService
	userService := user.UserService{}
	userInfo, err := userService.GetUserById(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": userInfo})
	return
}

// 添加好友
func AddFriends(c *gin.Context) {

}

// 获取好友列表
func GetUserList(c *gin.Context) {

}
