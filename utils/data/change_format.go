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
	}
	return
}

func ToInt32(i interface{}) (num int32, err error) {
	switch i.(type) {
	case string:
		n, e := strconv.Atoi(i.(string))
		num = int32(n)
		err = e
	}
	return
}

func ToInt64(i interface{}) (num int64, err error) {
	switch i.(type) {
	case int32:
		num = int64(i.(int32))
	case int64:
		num = i.(int64)
	case string:
		num, err = strconv.ParseInt(i.(string), 10, 64)
	}
	return
}

func ToFloat32(i interface{}) (num float32, err error) {
	return
}

func ToFloat64(i interface{}) (num float64, err error) {
	return
}
