package jwt

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"xiliangzi_pro/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {

		var token string
		if c.GetHeader("Authorization") != "" {
			token = c.GetHeader("Authorization")
		} else if c.Request.Header.Get("Sec-WebSocket-Protocol") != "" {
			token = c.Request.Header.Get("Sec-WebSocket-Protocol")
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "请求参数错误"})
			log.Fatal(gin.H{"msg": "请求参数错误"})
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token鉴权失败"})
			log.Fatal(gin.H{"msg": "Token鉴权失败"})
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token已过期"})
			log.Fatal(gin.H{"msg": "Token已过期"})
			c.Abort()
			return
		}
		// 将用户消息存储在context
		userMap := map[string]string{
			"uid":      strconv.Itoa(claims.Uid),
			"avatar":   claims.Avatar,
			"username": claims.Username,
			"nickname": claims.Nickname,
			"mobile":   claims.Mobile,
			"email":    claims.Email,
		}
		c.Set("user", userMap)
		c.Next()
	}
}
