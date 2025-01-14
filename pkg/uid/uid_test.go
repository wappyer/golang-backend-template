package uid

import (
	"fmt"
	"testing"
	"time"
)

func TestWorker_GetInt64(t *testing.T) {
	longUid := NewClient(
		8,                // 可拥有 2^8-1=255 个节点
		10,               // 每个周期生成id上限为 2^10-1=1023 个
		int64(1),         // 节点编号
		1293811200000,    // 2011-01-01 00:00:00 起始时间戳 控制在18位
		time.Millisecond, // 以1毫秒为一个周期
	)
	fmt.Printf("longUid: %d \n", longUid.GetInt64())

	shortUid := NewClient(
		4,                   // 可拥有 2^8-1=255 个节点
		4,                   // 每个周期生成id上限为 2^10-1=1023 个
		int64(1),            // 节点编号
		129381120000,        // 2011-01-01 00:00:00 起始时间戳 控制在14位
		time.Millisecond*10, // 以10毫秒为一个周期
	)
	fmt.Printf("shortUid: %d \n", shortUid.GetInt64())
}
