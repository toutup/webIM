package xiliangzi

import (
	"time"
)

// User [...]
type User struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`
	Username      string    `gorm:"column:username" json:"username"`
	Password      string    `gorm:"column:password" json:"password"`
	Email         string    `gorm:"column:email" json:"email"`
	Mobile        string    `gorm:"column:mobile" json:"mobile"`
	Nickname      string    `gorm:"column:nickname" json:"nickname"`
	Status        int       `gorm:"column:status" json:"status"`
	LastLoginTime time.Time `gorm:"column:last_login_time" json:"last_login_time"`
	CreateAt      time.Time `gorm:"column:createAt" json:"create_at"`
	UpdateAt      time.Time `gorm:"column:updateAt" json:"update_at"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}
