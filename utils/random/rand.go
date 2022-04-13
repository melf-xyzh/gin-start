/**
 * @Time    :2022/3/29 9:08
 * @Author  :ZhangXiaoyu
 */

package random

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	// CAPITAL 包含大写字母
	CAPITAL = 1
	// LOWERCASE 包含小写字母
	LOWERCASE = 2
	// SPECIAL 包含特殊字符
	SPECIAL = 4
	// NUMBER 包含数字
	NUMBER = 8
)

var (
	// 设置随机种子
	r = rand.New(rand.NewSource(time.Now().Unix()))
	// 大写字母
	capital *[]int
	// 小写字母
	lowercase *[]int
	// 特殊符号
	special *[]int
	// 数字
	number *[]int

	once sync.Once
)

// RandInt
/**
 *  @Description: 随机整数
 *  @param start
 *  @param end
 *  @return v
 */
func RandInt(min, max int) (v int) {
	return r.Intn(max-min) + min
}

// RandFloat
/**
 *  @Description: 随机小数
 *  @param min
 *  @param max
 *  @return v
 */
func RandFloat(min, max float64) (v float64) {
	return min + r.Float64()*(max-min)
}

// initASCII
/**
 *  @Description: 初始化ASCII码列表
 */
func initASCII() {
	once.Do(func() {
		fmt.Println("初始化列表")
		// 大写字母
		c := make([]int, 26)
		for i := 0; i < 26; i++ {
			c[i] = 65 + i
		}
		// 小写字母
		capital = &c
		l := make([]int, 26)
		for i := 0; i < 26; i++ {
			l[i] = 97 + i
		}
		lowercase = &l
		// 数字
		n := make([]int, 10)
		for i := 0; i < 10; i++ {
			n[i] = 48 + i
		}
		number = &n
		// 特殊字符(. @$!%*#_~?&^)
		s := []int{46, 64, 36, 33, 37, 42, 35, 95, 126, 63, 38, 94}
		special = &s
	})
	return
}

// RandString
/**
 *  @Description: 随机生成字符串
 *  @param n 字符串长度
 *  @param mode 字符串模式 random.NUMBER|random.LOWERCASE|random.SPECIAL|random.CAPITAL)
 *  @return str 生成的字符串
 */
func RandString(n int, mode int) (str string) {
	initASCII()
	var ascii []int
	if mode&CAPITAL >= CAPITAL {
		ascii = append(ascii, *capital...)
	}
	if mode&LOWERCASE >= LOWERCASE {
		ascii = append(ascii, *lowercase...)
	}
	if mode&SPECIAL >= SPECIAL {
		ascii = append(ascii, *special...)
	}
	if mode&NUMBER >= NUMBER {
		ascii = append(ascii, *number...)
	}
	if len(ascii) == 0 {
		return
	}
	var build strings.Builder
	for i := 0; i < n; i++ {
		build.WriteString(string(rune(ascii[r.Intn(len(ascii))])))
	}
	str = build.String()
	return
}
