/**
 * @Time    :2021/11/17 11:10
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:DatabaseForCasbin.go
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
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

var (
	DBForCasbin *gorm.DB
)

func InitDBForCasbin() {
	var err error
	// 获取配置文件
	config := Config.Config()


	// 拼接数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.DbUser, config.Database.DbPassword, config.Database.DbHost, config.Database.DbPort, config.Database.DbName,
	)

	DBForCasbin, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("初始化DB For Casbin失败")
		panic(err)
	}
}
