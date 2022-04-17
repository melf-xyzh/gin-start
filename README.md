# gin-start

### 介绍

一个Gin框架的项目公共基础库，避免进行基础模块的重复开发。

- 实现高并发可用的分布式雪花ID
- 封装跨域、限流等常用中间件
- 封装类型转换等常用函数
- 封装`Casbin`的常用方法
- 封装请求统一返回格式

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
	"github.com/melf-xyzh/gin-start/global/check"
	"github.com/melf-xyzh/gin-start/middleware"
	"github.com/melf-xyzh/gin-start/user/model"
	"github.com/melf-xyzh/gin-start/utils/result"
	"go-guide/guide/api/v1"
	"go-guide/guide/model"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	init := conf.Init{}
	// 初始化环境变量
	global.E = init.Env(global.ENV_DEV)
	// 初始化Viper（读取配置文件）
	global.V = init.Viper()
	// 初始化数据库连接池
	global.DB = init.Database()
	// 初始化雪花ID节点
	global.Node = init.Node()
	// 初始化Redis连接池
	global.RDB = init.Redis()
	// 初始化参数校验器
	global.Validate = init.Validate()
	// 初始化Casbin
	global.Enforcer = init.Casbin()

	r := gin.New()
	r.Use(middleware.Cors())

	err := global.DB.AutoMigrate(
		usermod.User{},
		guidemod.Index{},
		guidemod.WebSite{},
		guidemod.WebsiteGroup{},
	)
	if err != nil {
		panic("数据迁移失败")
	}

	r.GET("/", func(c *gin.Context) {

		user := usermod.User{
			Model: global.Model{
				ID:         global.CreateId(),
				CreateTime: global.CreateTime(),
			},
			Name:        "MELF",
			Email:       "123456789@99.com",
			LastLoginIp: "8.8.8.8",
		}
		err = check.Check(user, usermod.CreateUserCheck)
		if err != nil {
			return
		}
		//err = global.DB.Create(&user).Error

		var userFind usermod.User
		global.DB.First(&userFind)
		userFind.LastLoginIp = "192.168.1.11"
		global.DB.Updates(&userFind)
		result.OkDataMsg(c, userFind, "创建成功")
	})

	r.GET("/index", guideapi.IndexAPI.GetIndex)

	r.GET("/aaa/", middleware.Rate0("4-H"), func(c *gin.Context) {
		var userFind usermod.User
		global.DB.First(&userFind)
		result.OkDataMsg(c, userFind, "创建成功")
	})

	// 启动服务
	init.Run(r)
}
```

### 类型转换封装

```go
func main(){
    // var x interface{}
    
    // 将任意类型转换为string
	str := data.ToString(x)
	// 将任意类型转换为float32
	f32, err := data.ToFloat32(x)
	if err != nil {
		return
	}
	// 将任意类型转换为float64
	f64, err := data.ToFloat64(x)
	if err != nil {
		return
	}
	// 将任意类型转换为int
	i, err := data.ToInt(x)
	if err != nil {
		return
	}
	// 将任意类型转换为int32
	i32, err := data.ToInt32(x)
	if err != nil {
		return
	}
	// 将任意类型转换为int64
	i64, err := data.ToInt64(x)
	if err != nil {
		return
	}
}
```

### 中间件封装

```go
// 允许跨域
r.Use(middleware.Cors())

// 限流中间件
 *  5 reqs/second: "5-S"
 *  10 reqs/minute: "10-M"
 *  1000 reqs/hour: "1000-H"
 *  2000 reqs/day: "2000-D"
// 限流每个ip的访问频率
r.Use(middleware.Rate("4-H"))
// 限制该接口的总访问频率
r.Use(middleware.Rate0("4-H"))
```

### 非侵入式校验器

模型定义

```go
type User struct {
	global.Model
	Account       string     `json:"account"           gorm:"column:account;comment:账号;type:varchar(30)"`
	Password      string     `json:"-"                 gorm:"column:password;comment:密码;type:varchar(255)"`
	Name          string     `json:"name"              gorm:"column:name;comment:姓名;type:varchar(10)"`
	Avatar        string     `json:"avatar"            gorm:"column:avatar;comment:头像;type:varchar(255)"`
	Nickname      string     `json:"nickname"          gorm:"column:nickname;comment:昵称;type:varchar(20)"`
	Status        string     `json:"status"            gorm:"column:status;comment:状态;default:enable;type:varchar(10)"`
	UserType      string     `json:"userType"          gorm:"column:user_type;comment:用户类型;default:common;type:varchar(20)"`
	Phone         string     `json:"phone"             gorm:"column:phone;comment:手机号;type:varchar(255)"`
	Email         string     `json:"email"             gorm:"column:email;comment:邮箱;type:varchar(255)"`
	LoginCount    uint64     `json:"loginCount"        gorm:"column:login_count;comment:登录次数;default:0"`
	LastLoginTime *time.Time `json:"lastLoginTime"     gorm:"column:last_login_time;comment:上传登录时间;"`
	LastLoginIp   string     `json:"lastLoginIp"       gorm:"column:last_login_ip;comment:上次登录IP;type:varchar(100)"`
}
```

校验规则定义

```go
var (
	// CreateUserCheck 创建用户校验
	CreateUserCheck = check.Rules{
        // 键名为结构体字段名，规则请参考
        // https://github.com/go-playground/validator/blob/master/README.md
        "Name": "required",
	}
)
```

在这里向`https://github.com/go-playground/validator`致敬！

校验

```go
// 创建一个实例
user := usermod.User{
    Model: global.Model{
        ID:         global.CreateId(),
        CreateTime: global.CreateTime(),
    },
    Name:        "MELF",
    Email:       "123456789@99.com",
    LastLoginIp: "8.8.8.8",
}
// 校验该实例是否符合规则
err = check.Check(user, usermod.CreateUserCheck)
if err != nil {
    return
}
```

### 第三方接口请求（GET / POST方式）

```go
func main(){
	data, err := httpapi.Get.Get("https://api.uomg.com/api/comments.163")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(data)


	data, err = httpapi.Get.GetWithParams("https://api.uomg.com/api/comments.163", map[string]interface{}{"format":"text"})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(data)

	data,err = httpapi.Post.PostJson("需要请求的url",map[string]interface{}{"format":"text"})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(data)

	data,err = httpapi.Post.PostFormData("需要请求的url",map[string]interface{}{"format":"text"})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(data)

	data,err = httpapi.Post.PostUrlencoded("需要请求的url",map[string]interface{}{"format":"text"})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(data)
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
- [x] 跨域
- [ ] 身份认证（登录）
  - [ ] 基于Session
  - [ ] 基于Jwt
- [ ] 权限控制
  - [ ] 集成csabin
- [ ] SSE（服务端消息推送）
- [ ] gRPC
- [x] 分布式ID
  - [x] snowflow
- [ ] Websocket
- [x] 接口限流

### 备注

> codechina：https://codechina.csdn.net/qq_29537269/gin-start
>
> gitee：https://gitee.com/melf-xyzh/gin-start
>
> github：https://github.com/melf-xyzh/gin-start
