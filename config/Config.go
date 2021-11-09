/**
 * @Time    :2021/11/9 15:15
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:Config.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *		配置文件
 */

package config

import "time"

const (
	DbHost          = "127.0.0.1" // 数据库主机
	DbPort          = "3306"      // 数据库端口
	DbUser          = "root"      // 数据库用户名
	DbPassword      = "123456789" // 数据库密码
	DbName          = "gin-start" // 数据库名称
	MaxIdleConns    = 10          // 空闲连接池中连接的最大数量
	MaxOpenConns    = 100         // 打开数据库连接的最大数量
	ConnMaxLifetime = time.Hour   // 连接可复用的最大时间

	RouterHost = "0.0.0.0"
	RouterPort = "2048"
)
