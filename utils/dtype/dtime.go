/**
 * @Time    :2022/2/22 17:00
 * @Author  :ZhangXiaoyu
 */

package dtype

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"time"
)

type Time time.Time

// MarshalJSON 重写MarshalJSON方法
func (t Time) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	tStr := tTime.Format("2006-01-02 15:04:05")
	// 注意 json 字符串风格要求
	return []byte(fmt.Sprintf("\"%v\"", tStr)), nil
}

// Value 写入数据库之前，对数据做类型转换
func (t Time) Value() (driver.Value, error) {
	// DistributedId 转换成 int64 类型
	tTime := time.Time(t)
	return tTime, nil
}

// Scan 将数据库中取出的数据，赋值给目标类型
func (t *Time) Scan(v interface{}) error {
	switch v.(type) {
	case time.Time:
		*t = Time(v.(time.Time))
	default:
		val := reflect.ValueOf(v)
		typ := reflect.Indirect(val).Type()
		return errors.New(typ.Name() + "类型处理错误")
	}
	return nil
}

// CreateTime
/**
 *  @Description: 创建一个时间戳
 *  @return Time
 */
func CreateTime() Time {
	t := time.Now()
	tTime := Time(t)
	return tTime
}
