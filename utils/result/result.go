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
