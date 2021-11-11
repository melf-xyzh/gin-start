/**
 * @Time    :2021/11/10 16:44
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:LoginViews.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package views

import (
	url "gin-start/routers/urlmap"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginView
/**
 * @Description: 登录视图
 * @param c
 */
func LoginView(c *gin.Context) {
	c.HTML(http.StatusOK, url.TmplMap["Login"], gin.H{
		"title": "登录",
	})
}
