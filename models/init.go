/**
 * @Time    :2021/11/9 15:17
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:init.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package models

import (
	"errors"
	conn "gin-start/connection"
)

// DBAutoMigrate
/**
 * @Description: 自动迁移
 * @param DB 数据库连接池
 */
func DBAutoMigrate() (err error) {
	// 自动迁移
	err = conn.DB.AutoMigrate(
	// 实体
	)
	if err != nil {
		return errors.New("自动迁移失败：" + err.Error())
	}
	return nil
}

// InitModel
/**
 * @Description: 初始化数据表
 */
func InitModel() {

}
