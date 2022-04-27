package user

import (
	"errors"
	"sync"
	"xiliangzi_pro/internal/pkg/structs"
	"xiliangzi_pro/models/repo"
)

type UserService struct {
	Repo *repo.UserRepo
}

var (
	once        sync.Once
	userService *UserService
)

func GetUserServiceInstance() *UserService {
	once.Do(func() {
		userService = &UserService{
			Repo: repo.NewUserRepo(),
		}
	})
	return userService
}

// 获取用户信息
func (u *UserService) GetUserById(uid int) (structs.UserInfo, error) {
	userInfo, err := u.Repo.GetUserById(uid)
	if !err {
		return structs.UserInfo{}, errors.New("用户不存在")
	}
	return userInfo, nil
}
