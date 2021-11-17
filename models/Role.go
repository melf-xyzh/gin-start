/**
 * @Time    :2021/11/12 14:18
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:Role.go
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

type Role struct {
	ID        int64          `gorm:"primaryKey;comment:角色ID"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
	RoleName  string         `gorm:"type:varchar(20);not null;comment:角色名称"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = ID.CreateId()
	return
}

// TableName 自定义表名
func (Role) TableName() string {
	return "role"
}
