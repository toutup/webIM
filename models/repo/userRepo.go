package repo

import (
	"xiliangzi_pro/internal/pkg/structs"
	"xiliangzi_pro/models"
	"xiliangzi_pro/models/xiliangzi"
)

// UserRepo 操作user相关数据
type UserRepo struct{}

// 初始化UserRepo
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// 注册用户
func (u *UserRepo) Register(req structs.RegisterRequest) (xiliangzi.User, error) {

	user := xiliangzi.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Avatar:   req.Nickname,
	}
	if err := xiliangzi.UserMgr(models.GetDB()).Create(&user).Error; err != nil {
		return xiliangzi.User{}, err
	}
	return user, nil
}

// 用户登录
func (u *UserRepo) CheckAuth(userrname, password string) (structs.UserInfo, bool) {
	var login structs.UserInfo
	xiliangzi.UserMgr(models.GetDB()).Where(structs.LoginRequest{Username: userrname, Password: password}).First(&login)
	if login.Id > 0 {
		return login, true
	}
	return structs.UserInfo{}, false
}

// 获取用户信息
func (u *UserRepo) GetUserById(uid int) (structs.UserInfo, bool) {
	var userInfo structs.UserInfo
	xiliangzi.UserMgr(models.GetDB()).Where("id = ?", uid).First(&userInfo)
	if userInfo.Id > 0 {
		return userInfo, true
	}
	return structs.UserInfo{}, false
}
