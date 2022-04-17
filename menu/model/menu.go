/**
 * @Time    :2022/3/6 18:24
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:menu.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package model

import "github.com/melf-xyzh/gin-start/global"

type Menu struct {
	global.Model
	MenuName   string
	Sort       uint
	Url        string
	RouterName string
	IsShow     int
}
