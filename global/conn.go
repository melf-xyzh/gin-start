/**
 * @Time    :2022/2/19 7:45
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:conn.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *		单例连接池
 */

package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	E        *Env                // 环境类型实例
	V        *viper.Viper        // Viper实例
	Enforcer *casbin.Enforcer    // Casbin执行者
	DB       *gorm.DB            // 数据库连接池
	RDB      *redis.Client       // Redis连接池
	Validate *validator.Validate // Validate参数校验器
)
