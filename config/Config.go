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
	Self     SelfConfig     `json:"Self"`
}

type DatabaseConfig struct {
	Type            string `json:"Type"`
	DbHost          string `json:"DbHost"`
	DbPort          string `json:"DbPort"`
	DbUser          string `json:"DbUser"`
	DbPassword      string `json:"DbPassword"`
	DbName          string `json:"DbName"`
	MaxIdleConns    int    `json:"MaxIdleConns"`
	MaxOpenConns    int    `json:"MaxOpenConns"`
	ConnMaxLifetime int    `json:"ConnMaxLifetime"`
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
