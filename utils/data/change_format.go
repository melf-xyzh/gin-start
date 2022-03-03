/**
 * @Time    :2022/2/20 13:11
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:change_format.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package data

import (
	"fmt"
	"strconv"
	"time"
)

// ToString
/**
 * @Description: 转换为string类型
 * @param data
 * @return str
 */
func ToString(i interface{}) (str string) {
	switch i.(type) {
	case string:
		str = i.(string)
	case int:
		str = strconv.Itoa(i.(int))
	case int32:
		str = string(i.(int32))
	case int64:
		str = strconv.FormatInt(i.(int64), 10)
	case float32:
		str = fmt.Sprintf("%f", i.(float32))
	case float64:
		str = strconv.FormatFloat(i.(float64), 'f', -1, 32)
	case time.Time:
		str = i.(time.Time).Format("2006-01-02 15:04:05")
	case []uint8:
		var ba []byte
		for _, b := range i.([]uint8) {
			ba = append(ba, byte(b))
		}
		str = string(ba)
	case error:
		str = i.(error).Error()
	default:
		panic("该类型暂不支持")
	}
	return
}

func ToInt(i interface{}) (num int, err error) {
	switch i.(type) {
	case string:
		num, err = strconv.Atoi(i.(string))
	}
	return
}

// ToInt32
/**
 * @Description: 将任意类型转为int32类型
 * @param i
 * @return num
 * @return err
 */
func ToInt32(i interface{}) (num int32, err error) {
	switch i.(type) {
	case int:
		num = int32(i.(int))
	case int32:
		num = i.(int32)
	case int64:
		// 有可能造成精度丢失
		num = int32(i.(int64))
	case float32:
		// 有可能造成精度丢失
		num = int32(i.(float32))
	case float64:
		// 有可能造成精度丢失
		num = int32(i.(float64))
	case string:
		n, e := strconv.Atoi(i.(string))
		num = int32(n)
		err = e
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToInt64
/**
 * @Description: 将任意类型转为int64类型
 * @param i
 * @return num
 * @return err
 */
func ToInt64(i interface{}) (num int64, err error) {
	switch i.(type) {
	case int:
		num = int64(i.(int))
	case int32:
		num = int64(i.(int32))
	case int64:
		num = i.(int64)
	case float32:
		num = int64(i.(float32))
	case float64:
		num = int64(i.(float64))
	case string:
		num, err = strconv.ParseInt(i.(string), 10, 64)
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToFloat32
/**
 * @Description: 将任意类型转为float32类型
 * @param i
 * @return num
 * @return err
 */
func ToFloat32(i interface{}) (num float32, err error) {
	switch i.(type) {
	case string:
		// string无法直接转换float32，只能先转换为float64，再通过float64转float32
		var num64 float64
		num64, err = strconv.ParseFloat(i.(string), 32)
		num = float32(num64)
	case int:
		num = float32(i.(int))
	case int32:
		num = float32(i.(int32))
	case int64:
		num = float32(i.(int64))
	case float32:
		num = i.(float32)
	case float64:
		// 可能造成精度丢失
		num = float32(i.(float64))
	default:
		panic("该类型暂不支持")
	}
	return
}

// ToFloat64
/**
 * @Description: 将任意类型转为float64类型
 * @param i
 * @return num
 * @return err
 */
func ToFloat64(i interface{}) (num float64, err error) {
	switch i.(type) {
	case string:
		num, err = strconv.ParseFloat(i.(string), 64)
	case int:
		num = float64(i.(int))
	case int32:
		num = float64(i.(int32))
	case int64:
		num = float64(i.(int64))
	case float32:
		num = float64(i.(float32))
	case float64:
		num = i.(float64)
	default:
		panic("该类型暂不支持")
	}
	return
}
