/**
 * @Time    :2021/11/9 15:26
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:ViewRouters.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 * 		路由文件
 */

package routers

import (
	url "gin-start/routers/urlmap"
	"gin-start/views"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	// 主页
	r.GET(url.ViewRelativePath["Index"], views.IndexView)
	// 登录视图
	r.GET(url.ViewRelativePath["Login"], views.LoginView)
}
