# gin-start

### 介绍

一个Gin框架的项目公共基础库，避免进行基础模块的重复开发。

### 安装

```
go get github.com/melf-xyzh/gin-start
```

### 例子

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/melf-xyzh/gin-start/config"
	"github.com/melf-xyzh/gin-start/global"
	usermod "github.com/melf-xyzh/gin-start/user/model"
	"github.com/melf-xyzh/gin-start/utils/result"
	"log"
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

	err := global.DB.AutoMigrate(usermod.User{})
	if err != nil {
		panic("数据迁移失败")
	}

	r.GET("/", func(c *gin.Context) {
		user := usermod.User{
			Name: "MELF",
		}
		err := global.DB.Create(&user).Error
		log.Println(err)

		var userFind usermod.User
		global.DB.First(&userFind)
		userFind.LastLoginIp = "192.168.1.11"
		global.DB.Updates(&userFind)
		result.OkDataMsg(c, user, "创建成功")
	})

	// 启动服务
	init.Run(r)
}
```

### 项目目录

```bash
├─config            #配置相关
├─global            #全局变量（或单例连接池）、全局常量
├─resource          #资源文件夹
├─router            #路由文件夹
├─user              #用户模块
└─utils             #工具类
    ├─data          #数据格式相关
    ├─distributed #分布式相关
    └─result#返回消息封装
```

### 功能目标

- [x] Viper配置文件解析
- [x] 集成ORM（Gorm）
- [ ] 跨域
- [ ] 身份认证（登录）
  - [ ] 基于Session
  - [ ] 基于Jwt
- [ ] 权限控制
  - [ ] 集成csabin
- [ ] SSE（服务端消息推送）
- [ ] gRPC
- [ ] 分布式ID
  - [ ] snowflow
- [ ] Websocket
- [ ] 接口限流

### 备注

> codechina：https://codechina.csdn.net/qq_29537269/gin-start
>
> gitee：https://gitee.com/melf-xyzh/gin-start
>
> github：https://github.com/melf-xyzh/gin-start
