/**
 * @Time    :2021/11/11 11:06
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:Database.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package connection

import (
	"fmt"
	Config "gin-start/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

// InitDB
/**
 * @Description: 初始化数据库
 * @return err 错误
 */
func InitDB() (err error) {
	// 获取配置文件
	config := Config.Config()

	switch config.Database.Type {
	case "Mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.DbUser, config.Database.DbPassword, config.Database.DbHost, config.Database.DbPort, config.Database.DbName,
		)
		fmt.Println("连接数据库")
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
			//fmt.Println("数据库连接失败")
			panic("数据库连接失败")
		}
	default:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.DbUser, config.Database.DbPassword, config.Database.DbHost, config.Database.DbPort, config.Database.DbName,
		)
		fmt.Println("连接数据库")
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
			fmt.Println("数据库连接失败")
		}
	}

	// 维护连接池
	theDB, _ := DB.DB()
	// 设置空闲连接池中连接的最大数量
	theDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	// 设置打开数据库连接的最大数量
	theDB.SetMaxOpenConns(config.Database.MaxOpenConns)
	// 设置了连接可复用的最大时间
	theDB.SetConnMaxLifetime(time.Hour * time.Duration(config.Database.ConnMaxLifetime))

	//测试连通性
	return nil
}
