/**
 * @Time    :2022/3/30 16:15
 * @Author  :ZhangXiaoyu
 */

package mapstruct

import (
	"github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type User struct {
	Name        string   `json:"name"`
	PhoneNumber string   `json:"phone_number"`
	UserName    string   `json:"userName"`
	Hobbies     []string `json:"hobbies"`
}

// MapToStruct0
/**
 *  @Description: Map转Struct
 *  @param m
 *  @param data
 *  @return err
 */
func MapToStruct0(m map[string]interface{}, data interface{}) (err error) {
	err = mapstructure.Decode(m, &data)
	return
}

// MapToStruct
/**
 *  @Description: Map转Struct
 *  @param m
 *  @param data
 *  @return err
 */
func MapToStruct(m map[string]interface{}, data interface{}) (err error) {
	// 序列化
	arr, err := json.Marshal(m)
	if err != nil {
		return
	}
	// 反序列化
	err2 := json.Unmarshal(arr, &data)
	if err2 != nil {
		return
	}
	return
}

// StructToMap
/**
 *  @Description: Struct转Map
 *  @param obj
 *  @return map[string]interface{}
 */
func StructToMap(obj interface{}) map[string]interface{} {
	types := reflect.TypeOf(obj)
	values := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < types.NumField(); i++ {
		data[types.Field(i).Name] = values.Field(i).Interface()
	}
	return data
}
