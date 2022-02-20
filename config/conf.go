/**
 * @Time    :2022/2/19 9:36
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:conf.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package conf

import (
	"context"
	"database/sql"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/melf-xyzh/gin-start/global"
	"github.com/soheilhy/cmux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"strings"
	"time"
)

type Init struct{}

// Env
/**
 * @Description: 初始化环境
 * @receiver i
 * @param e
 * @return env
 */
func (i Init) Env(e global.Env) *global.Env {
	return &e
}

// Viper
/**
 * @Description:  初始化Viper(读取配置文件)
 * @receiver i
 * @return *viper.Viper
 */
func (i Init) Viper() *viper.Viper {
	v := viper.New()
	// 配置文件名
	v.SetConfigName(*global.E + ".config")
	// 配置文件类型
	v.SetConfigType("json")
	// 配置文件路径
	v.AddConfigPath("resource")
	// 读取配置文件信息
	err := v.ReadInConfig()
	if err != nil {
		panic("读取配置文件信息失败：" + err.Error())
	}
	return v
}

// Database
/**
 * @Description: 初始化数据库
 * @receiver i
 */
func (i Init) Database() (db *gorm.DB) {
	var err error
	dbType := global.V.GetString("Database.Type")
	dbHost := global.V.GetString("Database.DbHost")
	dbPort := global.V.GetString("Database.DbPort")
	dbUser := global.V.GetString("Database.DbUser")
	dbPassword := global.V.GetString("Database.DbPassword")
	dbName := global.V.GetString("Database.DbName")

	switch strings.ToLower(dbType) {
	case "mysql":
		dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
		mysqlConfig := mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}
		db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
		if err != nil {
			panic("初始化数据库连接池失败：" + err.Error())
		}
	default:
		panic("暂不支持")
	}

	maxIdleConns := global.V.GetInt("Database.MaxIdleConns")
	maxOpenConns := global.V.GetInt("Database.MaxOpenConns")
	connMaxLifetime := global.V.GetInt("Database.ConnMaxLifetime")
	connMaxIdleTime := global.V.GetInt("Database.ConnMaxIdleTime")
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxIdleConns)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(maxOpenConns)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(connMaxLifetime))
	// 连接池里面的连接最大空闲时长
	sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(connMaxIdleTime))
	return db
}

// Redis
/**
 * @Description: 初始化Redis
 * @receiver i
 */
func (i Init) Redis() *redis.Client {
	host := global.V.GetString("Redis.Host")
	port := global.V.GetString("Redis.Port")
	password := global.V.GetString("Redis.Password")
	db := global.V.GetInt("Redis.DB")
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	var ctx = context.Background()
	// 可连接性检测
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Redis连接池初始失败：" + err.Error())
	}
	return rdb
}

// Casbin
/**
 * @Description: 初始化Casbin
 * @receiver i
 */
func (i Init) Casbin() (Enforcer *casbin.Enforcer) {
	// Gorm适配器
	adapter, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		panic("Casbin Gorm适配器错误：" + err.Error())
	}
	log.Println("导入适配器")

	// 通过ORM新建一个执行者
	Enforcer, err = casbin.NewEnforcer("resource\\rbac_model.conf", adapter)
	if err != nil {
		panic("新建Casbin执行者异常：" + err.Error())
	}
	// 导入访问策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		panic("导入访问策略异常：" + err.Error())
	}
	return Enforcer
}

// Run
/**
 * @Description: 启动服务
 * @receiver i
 * @param r
 */
func (i Init) Run(r *gin.Engine) {
	port := global.V.GetString("Self.RouterPort")
	// 创建一个listener
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个cmux.
	m := cmux.New(l)
	go func() {
		httpL := m.Match(cmux.HTTP1Fast())
		err = r.RunListener(httpL)
		if err != nil {
			return
		}
	}()

	// 启动端口监听
	err = m.Serve()
	if err != nil {
		panic("服务启动失败：" + err.Error())
	}
}
