package xiliangzi

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取
func (obj *_UserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取
func (obj *_UserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithEmail email获取
func (obj *_UserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithMobile mobile获取
func (obj *_UserMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithNickname nickname获取
func (obj *_UserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithStatus status获取
func (obj *_UserMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithLastLoginTime last_login_time获取
func (obj *_UserMgr) WithLastLoginTime(lastLoginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_login_time"] = lastLoginTime })
}

// WithCreateAt createAt获取
func (obj *_UserMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["createAt"] = createAt })
}

// WithUpdateAt updateAt获取
func (obj *_UserMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updateAt"] = updateAt })
}

// WithAvatar avatar获取
func (obj *_UserMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UserMgr) GetFromID(id int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UserMgr) GetBatchFromID(ids []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容
func (obj *_UserMgr) GetFromUsername(username string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` = ?", username).First(&result).Error

	return
}

// GetBatchFromUsername 批量查找
func (obj *_UserMgr) GetBatchFromUsername(usernames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_UserMgr) GetFromPassword(password string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_UserMgr) GetBatchFromPassword(passwords []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_UserMgr) GetFromEmail(email string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`email` = ?", email).First(&result).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_UserMgr) GetBatchFromEmail(emails []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容
func (obj *_UserMgr) GetFromMobile(mobile string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`mobile` = ?", mobile).First(&result).Error

	return
}

// GetBatchFromMobile 批量查找
func (obj *_UserMgr) GetBatchFromMobile(mobiles []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容
func (obj *_UserMgr) GetFromNickname(nickname string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找
func (obj *_UserMgr) GetBatchFromNickname(nicknames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_UserMgr) GetFromStatus(status int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_UserMgr) GetBatchFromStatus(statuss []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromLastLoginTime 通过last_login_time获取内容
func (obj *_UserMgr) GetFromLastLoginTime(lastLoginTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`last_login_time` = ?", lastLoginTime).Find(&results).Error

	return
}

// GetBatchFromLastLoginTime 批量查找
func (obj *_UserMgr) GetBatchFromLastLoginTime(lastLoginTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`last_login_time` IN (?)", lastLoginTimes).Find(&results).Error

	return
}

// GetFromCreateAt 通过createAt获取内容
func (obj *_UserMgr) GetFromCreateAt(createAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`createAt` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找
func (obj *_UserMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`createAt` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过updateAt获取内容
func (obj *_UserMgr) GetFromUpdateAt(updateAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`updateAt` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找
func (obj *_UserMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`updateAt` IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容
func (obj *_UserMgr) GetFromAvatar(avatar string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找
func (obj *_UserMgr) GetBatchFromAvatar(avatars []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUserUsernameUIndex primary or index 获取唯一内容
func (obj *_UserMgr) FetchUniqueByUserUsernameUIndex(username string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` = ?", username).First(&result).Error

	return
}

// FetchUniqueByUserEmailUIndex primary or index 获取唯一内容
func (obj *_UserMgr) FetchUniqueByUserEmailUIndex(email string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`email` = ?", email).First(&result).Error

	return
}

// FetchUniqueByUserMobileUIndex primary or index 获取唯一内容
func (obj *_UserMgr) FetchUniqueByUserMobileUIndex(mobile string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`mobile` = ?", mobile).First(&result).Error

	return
}
