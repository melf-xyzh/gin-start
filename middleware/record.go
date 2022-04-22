/**
 * @Time    :2022/4/22 17:10
 * @Author  :ZhangXiaoyu
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/melf-xyzh/gin-start/global"
	"github.com/mssola/user_agent"
)

// AccessRecord 访问记录中间件
func AccessRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.UserAgent())
		ua := user_agent.New(c.Request.UserAgent())

		fmt.Printf("%v\n", ua.Mobile())  // => true
		fmt.Printf("%v\n", ua.Bot())     // => false
		fmt.Printf("%v\n", ua.Mozilla()) // => "5.0"

		fmt.Printf("%v\n", ua.Platform()) // => "Linux"
		fmt.Printf("%v\n", ua.OS())       // => "Android 2.3.7"

		name, version := ua.Engine()
		fmt.Printf("%v\n", name)    // => "AppleWebKit"
		fmt.Printf("%v\n", version) // => "533.1"

		name, version = ua.Browser()
		fmt.Printf("%v\n", name)    // => "Android"
		fmt.Printf("%v\n", version) // => "4.0"

		// Let's see an example with a bot.

		ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

		fmt.Printf("%v\n", ua.Bot()) // => true

		name, version = ua.Browser()
		fmt.Printf("%v\n", name)    // => Googlebot
		fmt.Printf("%v\n", version) // => 2.1
	}
}

type Record struct {
	global.Model
	Ip       string
	Method   string
	Path     string
	Status   string
	Latency  string
	Agent    string
	Error    string
	Body     string
	Response string
	Location string
	Province string
	City     string
	District string
	Isp      string
}
