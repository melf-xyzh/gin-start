/**
 * @Time    :2021/11/11 11:05
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:Redis.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package connection

import (
	Config "gin-start/config"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	RDB *redis.Client
)

func InitRedis() (err error) {
	// 获取配置文件
	config := Config.Config()

	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password, // no password set
		DB:       config.Redis.DB,       // use default DB
	})
	log.Println(RDB.PoolStats())

	return nil
}
