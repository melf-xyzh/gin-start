/**
 * @Time    :2022/2/20 13:03
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:model.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package global

import "github.com/melf-xyzh/gin-start/utils/dtype"

type Model struct {
	ID         dtype.DistributedId `json:"id,omitempty"            gorm:"column:id;primary_key;type:varchar(20)"`
	CreateTime dtype.Time          `json:"createTime,omitempty"    gorm:"column:create_time;comment:创建时间;"`
	UpdateTime *dtype.Time         `json:"updateTime,omitempty"    gorm:"column:update_time;comment:更新时间;"`
}
