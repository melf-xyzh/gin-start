/**
 * @Time    :2022/2/19 7:45
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:constants.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *		全局常量
 */

package global

// Env 类型（string类型的别名）
type Env = string

const (
	ENV_DEV Env = "dev" // 开发环境
	ENV_PRO Env = "pro" // 生产环境
	ENV_FAT Env = "fat" // 测试环境
)
