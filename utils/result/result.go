/**
 * @Time    :2022/2/20 16:59
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:result.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package result

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type code string

const (
	CODE_SUCCESS code = "0"
	CODE_FAIL    code = "1"
)

type Response struct {
	Code    code        `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Result
/**
 * @Description: Gin统一返回结果
 * @param code
 * @param data
 * @param message
 * @param c
 */
func Result(code code, data interface{}, message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

// OkMsg
/**
 * @Description: 返回带消息的成功
 * @param c
 * @param message
 */
func OkMsg(c *gin.Context, message string) {
	Result(CODE_SUCCESS, nil, message, c)
}

// OkDataMsg
/**
 * @Description: 返回带消息和数据的成功
 * @param c
 * @param data
 * @param message
 */
func OkDataMsg(c *gin.Context, data interface{}, message string) {
	Result(CODE_SUCCESS, data, message, c)
}

// FailMsg
/**
 * @Description: 返回带消息的失败
 * @param c
 * @param err
 */
func FailMsg(c *gin.Context, err error) {
	Result(CODE_FAIL, nil, err.Error(), c)
}

// FailDataMsg
/**
 * @Description: 返回带消息和数据的失败
 * @param c
 * @param data
 * @param err
 */
func FailDataMsg(c *gin.Context, data interface{}, err error) {
	Result(CODE_FAIL, data, err.Error(), c)
}

// File
/**
 *  @Description: 返回文件
 *  @param c
 *  @param fileName
 *  @param filePath
 */
func File(c *gin.Context, fileName, filePath string) {
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
}

// OkView
/**
 * @Description: 返回成功界面
 * @param c
 * @param tmplPath tmpl文件路径
 * @param title 网页标题
 * @param msg
 */
func OkView(c *gin.Context, tmplPath, title, msg string) {
	c.HTML(http.StatusOK, tmplPath, gin.H{
		"code":  CODE_SUCCESS,
		"title": title,
		"msg":   msg,
	})
}

// FailView
/**
 * @Description: 返回失败界面
 * @param c
 * @param tmplPath tmpl文件路径
 * @param title 网页标题
 * @param msg
 */
func FailView(c *gin.Context, tmplPath, title, msg string) {
	c.HTML(http.StatusOK, tmplPath, gin.H{
		"code":  CODE_FAIL,
		"title": title,
		"msg":   msg,
	})
}

// View
/**
 * @Description: 返回界面
 * @param c
 * @param tmplPath
 * @param title
 */
func View(c *gin.Context, tmplPath, title string) {
	c.HTML(http.StatusOK, tmplPath, gin.H{
		"code":  CODE_FAIL,
		"title": title,
	})
}

// DataView
/**
 * @Description: 返回数据界面
 * @param c
 * @param tmplPath
 * @param title
 * @param data
 */
func DataView(c *gin.Context, tmplPath, title string, data map[string]interface{}) {
	data["title"] = title
	c.HTML(http.StatusOK, tmplPath, data)
}
