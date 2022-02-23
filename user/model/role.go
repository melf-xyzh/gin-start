/**
 * @Time    :2022/2/23 16:00
 * @Author  :ZhangXiaoyu
 */

package usermod

import (
	"github.com/melf-xyzh/gin-start/global"
	"github.com/melf-xyzh/gin-start/global/check"
	"gorm.io/gorm"
)

var (
	// CreateRoleCheck 创建角色校验
	CreateRoleCheck = check.Rules{
		"RoleName": "required",
	}
)

// Role 角色
type Role struct {
	global.Model
	RoleName string `json:"roleName"           gorm:"column:role_name;comment:角色名称;type:varchar(30)"`
	Remark   string `json:"remark"             gorm:"column:remark;comment:角色描述;type:varchar(255)"`
}

// TableName 自定义表名
func (Role) TableName() string {
	return "role"
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = global.CreateId()
	}
	r.CreateTime = global.CreateTime()
	return

}

func (r *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	t := global.CreateTime()
	r.UpdateTime = &t
	return
}
