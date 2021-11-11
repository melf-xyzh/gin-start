/**
 * @Time    :2021/11/11 14:55
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:IndexViews.go
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

// IndexView
/**
 * @Description: 主页
 * @param c
 */
func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, url.TmplMap["Index"], gin.H{
		"title": "主页",
	})
}
