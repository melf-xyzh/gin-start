/**
 * @Time    :2022/2/23 12:01
 * @Author  :ZhangXiaoyu
 */

package check

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/melf-xyzh/gin-start/global"
	"log"
	"reflect"
)

// 参考文档：https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

type Rules map[string]string

// Check
/**
 *  @Description: 对结构体内置数据类型根据validator库进行校验
 * （不支持slice、array、struct、map、channel、interface、function类型）
 *  @param s
 *  @param rules
 *  @return err
 */
func Check(s interface{}, rules Rules) (err error) {
	if rules == nil {
		return
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	k := v.Kind() // 获取到st对应的类别
	if k != reflect.Struct {
		log.Println(k)
	}

	// 遍历结构体的所有字段
	for i := 0; i < v.NumField(); i++ {
		tagVal := t.Field(i)
		val := v.Field(i)
		rule, ok := rules[tagVal.Name]
		if ok {
			err = CheckValue(val, rule)
			if err != nil {
				return err
			}
		}
	}
	return
}

// CheckStruct
/**
 *  @Description: 校验结构体
 *  @param s 结构体
 *  @return errF 错误
 */
func CheckStruct(s interface{}) (errF validator.FieldError) {
	err := global.Validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
		for _, errF = range err.(validator.ValidationErrors) {
			return errF
		}
	}
	return
}

// CheckField
/**
 *  @Description: 字段校验
 *  @param field 字段实例（"joeybloggs.gmail.com"）
 *  @param tag 校验标签（"required,email"）
 *  @return errs 返回错误
 */
func CheckField(field interface{}, tag string) (errs error) {
	errs = global.Validate.Var(field, tag)
	return
}

// CheckValue
/**
 *  @Description: 字段值校验
 *  @param v
 *  @param rule
 *  @return err
 */
func CheckValue(v reflect.Value, rule string) (err error) {
	switch v.Kind() {
	case reflect.String:
		err = global.Validate.Var(v.String(), rule)
	case reflect.Bool:
		err = global.Validate.Var(v.Bool(), rule)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		err = global.Validate.Var(v.Uint(), rule)
	case reflect.Float32, reflect.Float64:
		err = global.Validate.Var(v.Float(), rule)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		err = global.Validate.Var(v.Int(), rule)
	default:
		err = errors.New("暂不支持")
	}
	return err
}
