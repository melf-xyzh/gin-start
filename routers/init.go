/**
 * @Time    :2021/11/11 14:03
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:init.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package routers

import (
	"fmt"
	Config "gin-start/config"
	"gin-start/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 获取配置文件
	config := Config.Config()

	r := gin.Default()

	store, _ := redis.NewStore(10, "tcp", config.Redis.Host+":"+config.Redis.Port, config.Redis.Password, []byte("secret"))

	// 使用中间件
	r.Use(
		// 中间件
		sessions.Sessions("mysession", store),
		// session登录验证中间件
		middleware.AuthSessionMiddleWare(),
	)

	// 配置模板
	r.LoadHTMLGlob("./templates/**/*")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	r.StaticFS("/static", http.Dir("./static"))

	// 加载View的路由
	LoadRouter(r)

	err := r.Run(fmt.Sprintf("%s:%s", config.Self.Host, config.Self.Port))
	if err != nil {
		log.Fatal("初始化路由出错")
	}
	return r
}
