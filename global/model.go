/**
 * @Time    :2022/2/20 13:03
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:model.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package global

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sync"
	"time"

	"github.com/melf-xyzh/gin-start/utils/dtype"
)

var (
	// 定义一个锁
	idLock sync.Mutex
)

type Model struct {
	ID         dtype.DistributedId `json:"id,omitempty"            gorm:"column:id;primary_key;"`
	CreateTime dtype.Time          `json:"createTime,omitempty"    gorm:"column:create_time;comment:创建时间;"`
	UpdateTime *dtype.Time         `json:"updateTime,omitempty"    gorm:"column:update_time;comment:更新时间;"`
}

// CreateId
/**
 *  @Description: 创建一个分布式ID（雪花ID）
 *  @return DistributedId
 */
func CreateId() dtype.DistributedId {
	idLock.Lock()
	defer idLock.Unlock()

	// Create a new Node with a Node number of 1
	// node, err := snowflake.NewNode(nodeNum)
	// 为编号为nodeNum的节点生成一个节点
	nodeNum := V.GetInt64("Distributed.Node")
	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		fmt.Println(err)
	}
	id := node.Generate()
	return dtype.DistributedId(id.Int64())
}

// CreateTime
/**
 *  @Description: 创建一个时间戳
 *  @return Time
 */
func CreateTime() dtype.Time {
	t := time.Now()
	tTime := dtype.Time(t)
	return tTime
}
