/**
 * @Time    :2021/11/11 11:21
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:AuthSessionMiddleWare.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package middleware

import (
	url "gin-start/routers/urlmap"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// AuthSessionMiddleWare
/**
 * @Description: 身份认证中间件
 * @return gin.HandlerFunc
 */
func AuthSessionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.RequestURI)
		if strings.Contains(c.Request.RequestURI, "Login") {
			log.Println("不需要登录验证")
		} else if strings.Contains(c.Request.RequestURI, "static") {
			log.Println("不需要登录验证")
		} else if strings.Contains(c.Request.RequestURI, "api") {
			log.Println("不需要登录验证")
		} else {
			log.Println("登录验证")
			session := sessions.Default(c)
			userId := session.Get("userId")
			// 获取用户的userID
			log.Println(userId)
			if userId == nil {
				log.Println("未登录")
				// 重定向至登录界面
				c.Redirect(http.StatusMovedPermanently, url.ViewRelativePath["Login"])
			}

			isLogin := session.Get("isLogin")
			if isLogin != true {
				log.Println("未登录")
				// 重定向至登录界面
				c.Redirect(http.StatusMovedPermanently, url.ViewRelativePath["Login"])
			}
		}
	}
}
