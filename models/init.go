/**
 * @Time    :2021/11/9 15:17
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:init.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package models

import (
	"fmt"
	Config "gin-start/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DbUser, Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName,
	)
	fmt.Println("连接数据库")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库连接失败")
	}

	// 维护连接池
	theDB, _ := DB.DB()
	// 设置空闲连接池中连接的最大数量
	theDB.SetMaxIdleConns(Config.MaxIdleConns)
	// 设置打开数据库连接的最大数量
	theDB.SetMaxOpenConns(Config.MaxOpenConns)
	// 设置了连接可复用的最大时间
	theDB.SetConnMaxLifetime(Config.ConnMaxLifetime)

	// 自动迁移
	err = DB.AutoMigrate(
	// 实体
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("自动迁移失败")
	}

	// 初始化数据表

	//测试连通性
	return nil
}
