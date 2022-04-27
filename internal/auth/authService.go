package auth

import (
	"sync"
	"xiliangzi_pro/internal/pkg/structs"
	"xiliangzi_pro/models/repo"
	"xiliangzi_pro/models/xiliangzi"
	"xiliangzi_pro/pkg/util"
)

type AuthService struct {
	Repo *repo.UserRepo
}

var once sync.Once
var authService *AuthService

// 获取authService单例
func GetAuthServiceInstance() *AuthService {
	once.Do(func() {
		authService = &AuthService{
			Repo: repo.NewUserRepo(),
		}
	})
	return authService
}

// 注册用户
func (a *AuthService) Register(req structs.RegisterRequest) (structs.UserInfo, error) {
	// 处理password字段，对password做MD5
	pwd, err := util.GetMd5String(req.Password)
	if err != nil {
		return structs.UserInfo{}, err
	}
	req.Password = pwd
	user, err := a.Repo.Register(req)
	if err != nil {
		return structs.UserInfo{}, err
	}
	// 获取token
	token, err := util.GenerateToken(user)
	if err != nil {
		return structs.UserInfo{}, err
	}
	userInfo := structs.UserInfo{
		Id:       user.ID,
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Token:    token,
	}
	return userInfo, nil
}

// 用户登录
func (a *AuthService) Login(username, password string) (structs.UserInfo, error) {

	pwd, err := util.GetMd5String(password)
	if err != nil {
		return structs.UserInfo{}, err
	}
	// 验证用户是否存在
	userInfo, isExist := a.Repo.CheckAuth(username, pwd)
	if isExist {
		// 获取token

		token, err := util.GenerateToken(xiliangzi.User{
			ID:       userInfo.Id,
			Avatar:   userInfo.Avatar,
			Username: userInfo.Username,
			Nickname: userInfo.Nickname,
			Mobile:   userInfo.Mobile,
			Email:    userInfo.Email,
		})
		if err != nil {
			return structs.UserInfo{}, err
		}
		userInfo.Token = token
	}
	return userInfo, nil
}
