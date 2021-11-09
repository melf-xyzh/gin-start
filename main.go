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
	myDB "gin-start/models"
	"gin-start/routers"
)

func main() {
	// 连接数据库
	errDB := myDB.InitDB()
	if errDB != nil {
		fmt.Println("数据库异常")
		//panic(err)
	}

	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("服务启动失败, 异常：%v\n", err)
	}
}
