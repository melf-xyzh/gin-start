/**
 * @Time    :2022/2/22 16:59
 * @Author  :ZhangXiaoyu
 */

package dtype

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/melf-xyzh/gin-start/utils/data"
	"reflect"
	"strconv"
)

type DistributedId int64

// MarshalJSON 重写MarshalJSON方法
func (t DistributedId) MarshalJSON() ([]byte, error) {
	str := strconv.FormatInt(int64(t), 10)
	// 注意 json 字符串风格要求
	return []byte(fmt.Sprintf("\"%v\"", str)), nil
}

// Value 写入数据库之前，对数据做类型转换
func (t DistributedId) Value() (driver.Value, error) {
	// DistributedId 转换成 int64 类型
	num := int64(t)
	return num, nil
}

// Scan 将数据库中取出的数据，赋值给目标类型
func (t *DistributedId) Scan(v interface{}) error {
	switch v.(type) {
	case []uint8:
		numStr := data.ToString(v.([]uint8))
		num, _ := strconv.ParseInt(numStr, 10, 64)
		*t = DistributedId(num)
	default:
		val := reflect.ValueOf(v)
		typ := reflect.Indirect(val).Type()
		return errors.New(typ.Name() + "类型处理错误")
	}
	return nil
}
