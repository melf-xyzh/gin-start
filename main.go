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
	"fmt"
	"gin-start/config"
	"gin-start/connection"
	"gin-start/models"
	"gin-start/routers"
)

func main() {
	//初始化配置文件
	err := config.InitConfigJson("config\\Config.json")
	if err != nil {
		panic(err)
	}

	// 连接数据库
	errPool := connection.InitConnectionPool()
	if errPool != nil {
		panic(errPool)
	}

	// 数据库自动迁移
	errDBCreate := models.DBAutoMigrate()
	if errDBCreate != nil {
		panic(errDBCreate)
	}

	// 数据表数据初始化
	models.InitModel()


	// 初始化路由
	r := routers.InitRouter()
	if errRouter := r.Run(); errRouter != nil {
		fmt.Printf("服务启动失败, 异常：%v\n", errRouter)
	}
}
