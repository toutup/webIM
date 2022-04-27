package util

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

// GetMd5String 获取md5加密字符串
func GetMd5String(str string) (string, error) {

	//初始化hash.Hash接口
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return "", err
	}
	sum := md5.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx *gin.Context) map[string]string {
	return ctx.GetStringMapString("user")
}
