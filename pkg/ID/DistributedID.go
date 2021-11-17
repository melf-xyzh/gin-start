/**
 * @Time    :2021/11/12 11:11
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:DistributedID.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package ID

import (
	"fmt"
	Config "gin-start/config"
	"github.com/bwmarrin/snowflake"
)

func CreateId() int64 {
	// 获取配置文件
	config := Config.Config()
	node, err := snowflake.NewNode(config.Company.ID)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	id := node.Generate().Int64()
	fmt.Println("id is:", id)
	return id
}
