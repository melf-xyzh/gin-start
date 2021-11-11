/**
 * @Time    :2021/11/9 15:15
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:Config.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *		配置文件
 */

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type GlobalConfig struct {
	Database DatabaseConfig `json:"Database"`
	Redis    RedisConfig    `json:"Redis"`
	Self     SelfConfig     `json:"Self"`
}

type DatabaseConfig struct {
	Type            string `json:"Type"`            // 数据库类型
	DbHost          string `json:"DbHost"`          // 主机
	DbPort          string `json:"DbPort"`          // 端口
	DbUser          string `json:"DbUser"`          // 用户
	DbPassword      string `json:"DbPassword"`      // 密码
	DbName          string `json:"DbName"`          // 库名
	MaxIdleConns    int    `json:"MaxIdleConns"`    // 空闲连接池中连接的最大数量
	MaxOpenConns    int    `json:"MaxOpenConns"`    // 打开数据库连接的最大数量
	ConnMaxLifetime int    `json:"ConnMaxLifetime"` // 连接可复用的最大时间（小时）
}

type RedisConfig struct {
	Host     string `json:"Host"`     // 主机
	Port     string `json:"Port"`     // 端口
	Password string `json:"Password"` // 密码
	DB       int    `json:"DB"`       // 数据库（0-15）
}

type SelfConfig struct {
	Host string `json:"RouterHost"`
	Port string `json:"RouterPort"`
}

var (
	globalConfig *GlobalConfig
	configMux    sync.RWMutex
)

func Config() *GlobalConfig {
	return globalConfig
}

// InitConfigJson
/**
 * @Description: 导入配置文件
 * @param fliepath config文件的路径
 * @return error 错误
 */
func InitConfigJson(fliepath string) error {
	var config GlobalConfig
	file, err := ioutil.ReadFile(fliepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}

	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println("配置文件读取失败", err)
		return err
	}

	configMux.Lock()
	globalConfig = &config
	configMux.Unlock()
	return nil
}
