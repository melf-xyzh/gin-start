/**
 * @Time    :2022/2/20 13:20
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:id.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package distributed

import (
	"github.com/bwmarrin/snowflake"
	"github.com/melf-xyzh/gin-start/global"
)

func CreateId() (id int64) {
	// 为编号为nodeNum的节点生成一个节点
	nodeNum := global.V.GetInt64("Distributed.Node")
	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		panic("雪花ID生成失败：" + err.Error())
	}
	// 生成一个雪花ID
	id = node.Generate().Int64()
	return
}

func CreateIdString() (id string) {
	// 为编号为nodeNum的节点生成一个节点
	nodeNum := global.V.GetInt64("Distributed.Node")
	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		panic("雪花ID生成失败：" + err.Error())
	}
	// 生成一个雪花ID
	id = node.Generate().String()
	return
}
