/**
 * @Time    :2021/11/9 15:02
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:main.go
 * @Project :gin-start
 * @Software:GoLand
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package main

import (
	"gin-start/config"
	"gin-start/global"
	"github.com/gin-gonic/gin"
)

func main() {
	init := conf.Init{}
	// 初始化环境变量
	global.E = init.Env(global.ENV_DEV)
	// 初始化Viper（读取配置文件）
	global.V = init.Viper()
	// 初始化数据库连接池
	global.DB = init.Database()
	// 初始化Redis连接池
	global.RDB = init.Redis()

	r := gin.New()
	// 启动服务
	init.Run(r)
}
