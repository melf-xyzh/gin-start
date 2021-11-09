/**
 * @Time    :2021/11/9 15:26
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:Routers.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 * 		路由文件
 */

package routers

import (
	"fmt"
	Config "gin-start/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Init() *gin.Engine {
	// 获取配置文件
	config := Config.Config()

	r := gin.Default()

	// 使用中间件
	r.Use(
	// 中间件
	)

	// 配置模板
	r.LoadHTMLGlob("./templates/**/*")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	r.StaticFS("/static", http.Dir("./static"))

	err := r.Run(fmt.Sprintf("%s:%s", config.Self.Host, config.Self.Port))
	if err != nil {
		log.Fatal("初始化路由出错")
	}
	return r
}
