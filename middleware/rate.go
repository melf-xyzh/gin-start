/**
 * @Time    :2022/2/24 9:03
 * @Author  :ZhangXiaoyu
 */

package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/melf-xyzh/gin-start/global"
	"github.com/melf-xyzh/gin-start/utils/result"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"strconv"
)

// Rate
/**
 *  @Description:限流中间件（对每个client限流）
 *  @param formatted 限流规则
 *  5 reqs/second: "5-S"
 *  10 reqs/minute: "10-M"
 *  1000 reqs/hour: "1000-H"
 *  2000 reqs/day: "2000-D"
 *  @return gin.HandlerFunc
 */
// Rate 限流中间件
func Rate(formatted string) gin.HandlerFunc {
	l := newLimiter(formatted)
	return func(c *gin.Context) {
		// Client限流
		context, err := l.Get(c, c.ClientIP()+":"+c.Request.RequestURI)
		if err != nil {
			result.FailMsg(c, err)
			c.Abort()
			return
		}
		c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))
		if context.Reached {
			result.FailMsg(c, errors.New("接口访问超过限制"))
			c.Abort()
			return
		}
		c.Next()
	}
}

// Rate0 限流中间件（总访问量限流）
/**
 *  @Description:限流中间件
 *  @param formatted 限流规则
 *  5 reqs/second: "5-S"
 *  10 reqs/minute: "10-M"
 *  1000 reqs/hour: "1000-H"
 *  2000 reqs/day: "2000-D"
 *  @return gin.HandlerFunc
 */
func Rate0(formatted string) gin.HandlerFunc {
	l := newLimiter(formatted)
	return func(c *gin.Context) {
		// Client限流
		context, err := l.Get(c, "ALL:"+c.Request.RequestURI)
		if err != nil {
			result.FailMsg(c, err)
			c.Abort()
			return
		}
		c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))
		if context.Reached {
			result.FailMsg(c, errors.New("接口访问超过限制"))
			c.Abort()
			return
		}
		c.Next()
	}
}

// newLimiter
/**
 *  @Description: 创建一个限流器
 *  @param formatted
 *  @return l
 *  @return err
 */
func newLimiter(formatted string) (l *limiter.Limiter) {
	// 初始化速率
	rate, err := limiter.NewRateFromFormatted(formatted)
	if err != nil {
		panic(err)
	}
	// 初始化限流器存储
	var store limiter.Store
	// 优先使用redis
	if global.RDB != nil {
		store, err = redis.NewStoreWithOptions(global.RDB, limiter.StoreOptions{
			Prefix: "Rate",
		})
		if err != nil {
			panic("限流中间件(Redis)初始化出错:" + err.Error())
		}
	}
	// 若redis不可用，则使用内存
	if store == nil {
		store = memory.NewStore()
	}
	// 创建一个限流器
	l = limiter.New(store, rate)
	return
}
