/**
 * @Time    :2022/2/20 12:53
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:user.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package usermod

import (
	"github.com/melf-xyzh/gin-start/global"
	"github.com/melf-xyzh/gin-start/global/check"
	"gorm.io/gorm"
	"time"
)

var (
	// CreateUserCheck 创建用户校验
	CreateUserCheck = check.Rules{
		"Name": "required",
	}

	// UserStatus 用户状态
	UserStatus = map[string]string{
		"enable":  "enable",  // 启用
		"disable": "disable", // 禁用
	}
	// UserType 用户类型
	UserType = map[string]string{
		"admin":  "admin",  // 管理员
		"common": "common", // 普通用户
	}
)

type User struct {
	global.Model
	Account       string     `json:"account"           gorm:"column:account;comment:账号;type:varchar(30)"`
	Password      string     `json:"-"                 gorm:"column:password;comment:密码;type:varchar(255)"`
	Name          string     `json:"name"              gorm:"column:name;comment:姓名;type:varchar(10)"`
	Avatar        string     `json:"avatar"            gorm:"column:avatar;comment:头像;type:varchar(255)"`
	Nickname      string     `json:"nickname"          gorm:"column:nickname;comment:昵称;type:varchar(20)"`
	Status        string     `json:"status"            gorm:"column:status;comment:状态;default:enable;type:varchar(10)"`
	UserType      string     `json:"userType"          gorm:"column:user_type;comment:用户类型;default:common;type:varchar(20)"`
	Phone         string     `json:"phone"             gorm:"column:phone;comment:手机号;type:varchar(255)"`
	Email         string     `json:"email"             gorm:"column:email;comment:邮箱;type:varchar(255)"`
	LoginCount    uint64     `json:"loginCount"        gorm:"column:login_count;comment:登录次数;default:0"`
	LastLoginTime *time.Time `json:"lastLoginTime"     gorm:"column:last_login_time;comment:上传登录时间;"`
	LastLoginIp   string     `json:"lastLoginIp"       gorm:"column:last_login_ip;comment:上次登录IP;type:varchar(100)"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.ID = global.CreateId()
	}
	u.CreateTime = global.CreateTime()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	t := global.CreateTime()
	u.UpdateTime = &t
	return
}
