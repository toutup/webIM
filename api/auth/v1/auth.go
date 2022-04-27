package v1

import (
	"fmt"
	"net/http"
	"xiliangzi_pro/internal/auth"
	"xiliangzi_pro/internal/pkg/structs"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	// 接收参数
	var request structs.RegisterRequest
	// 解析JSON参数
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// 参数验证
	valid := validation.Validation{}
	valid.MinSize(request.Password, 6, "password").Message("不小于6位")
	valid.Email(request.Email, "email").Message("格式不正确")
	valid.Mobile(request.Mobile, "mobile").Message("格式不正确")
	valid.MaxSize(request.Username, 32, "username").Message("最大为32个字符")
	valid.MaxSize(request.Nickname, 32, "username").Message("最大为32个字符")
	if valid.HasErrors() {

		for _, err := range valid.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s:%s", err.Key, err.Message)})
			return
		}
	}
	// 初始化authService
	authService := auth.GetAuthServiceInstance()
	// 注册用户
	userInfo, err := authService.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": userInfo})
	return
}

func Login(c *gin.Context) {

	// 接收参数
	var request structs.LoginRequest
	// 解析JSON参数
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// 参数验证
	valid := validation.Validation{}
	valid.MinSize(request.Password, 6, "password").Message("不小于6位")
	valid.MaxSize(request.Username, 32, "username").Message("最大为32个字符")
	if valid.HasErrors() {

		for _, err := range valid.Errors {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s:%s", err.Key, err.Message)})
			return
		}
	}
	// 初始化authService
	authService := auth.GetAuthServiceInstance()
	// 验证用户
	userInfo, err := authService.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": userInfo})
	return
}
