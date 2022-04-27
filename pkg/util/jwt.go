package util

import (
	"os"
	"time"
	"xiliangzi_pro/models/xiliangzi"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWTSTRING"))

type Claims struct {
	Uid      int
	Avatar   string
	Username string
	Nickname string
	Email    string
	Mobile   string
	jwt.StandardClaims
}

// 生成token
func GenerateToken(user xiliangzi.User) (string, error) {

	nowTime := time.Now()
	exprieTime := nowTime.Add(24 * time.Hour)

	claims := &Claims{
		Uid:      user.ID,
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Mobile:   user.Mobile,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exprieTime.Unix(),
			Issuer:    "xiliangzi",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
