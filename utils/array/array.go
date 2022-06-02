/**
 * @Time    :2022/5/18 17:27
 * @Author  :Xiaoyu.Zhang
 */

package array

import (
	"errors"

)

func Unique(a interface{}) (answer interface{}, err error) {
	switch a.(type) {
	case []string:
		answer := make([]string, 0)
		m := make(map[string]bool) // map的值不重要
		for _, v := range a.([]string) {
			if _, ok := m[v]; !ok {
				answer = append(answer, v)
				m[v] = true
			}
		}
		return answer, nil
	case []int32:
		answer := make([]int32, 0)
		m := make(map[int32]bool) // map的值不重要
		for _, v := range a.([]int32) {
			if _, ok := m[v]; !ok {
				answer = append(answer, v)
				m[v] = true
			}
		}
		return answer, nil
	default:
		err = errors.New("不支持的类型")
		return
	}
	return
}

