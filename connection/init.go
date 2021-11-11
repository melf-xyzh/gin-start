/**
 * @Time    :2021/11/11 14:26
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:init.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package connection

import "errors"

// InitConnectionPool
/**
 * @Description: 初始化连接池
 * @return err 异常
 */

func InitConnectionPool() (err error) {
	// 初始化数据库连接池
	errDB := InitDB()
	if errDB != nil {
		return errors.New("数据库连接异常：" + errDB.Error())
	}
	// 初始化Redis连接池
	errRDB := InitRedis()
	if errRDB != nil {
		return errors.New("Redis连接异常：" + errRDB.Error())
	}
	return nil
}
