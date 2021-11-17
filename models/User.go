/**
 * @Time    :2021/11/12 14:04
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:User.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package models

import (
	"gin-start/pkg/ID"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          int64          `gorm:"primaryKey;comment:用户ID"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:删除时间"`
	Name        string         `gorm:"type:varchar(10);not null;comment:姓名"`
	Account     string         `gorm:"type:varchar(20);not null;comment:账号"`
	Password    string         `gorm:"type:varchar(255);not null;comment:密码"`
	Phone       string         `gorm:"type:varchar(255);not null;comment:电话"`
	Sex         int8           `gorm:"default:1;comment:性别"`
	LastLoginAt *time.Time     `gorm:"comment:上传登录时间"`
	LastLoginIP *string        `gorm:"comment:上传登录IP"`
	IsFirst     bool           `gorm:"default:false;comment:是否初次登录"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = ID.CreateId()
	return
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}
