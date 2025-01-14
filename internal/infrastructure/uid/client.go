package uid

import (
	"gitee.com/wappyer/golang-backend-template/pkg/uid"
	"sync"
	"time"
)

/*
 * 使用雪花算法的id生成器
 * 当节点数越多，周期越短，每个周期生成的id上限越高，对应的id位数也会越多
 * timeEpoch 设置在2011年是为了控制生成的id位数一致，不然起始的位数会比较少
 *
 * longUid  2131年前位数在18位；最多可有255个节点；1毫秒为一周期，每个周期生成的id上限为1023个；用于生成requestId等高频场景
 *
 * shortUid 2134年前位数在14位；最多可有15个节点；10毫秒为一周期，每个周期生成的id上限为15个；用于生成orderNo等低频场景
 *
 * miniUid 2134年前位数在10位；最多可有4个节点；100秒为一周期，每个周期生成的id上限为64个；用于生成条码对位数限制高场景
 * 注意：miniUid周期为100秒，如果100秒内生成的id数超过64个，那将会阻塞100秒！！！
 */

var longUid *uid.Client
var longUidOnce sync.Once

func LongUid(index int64) *uid.Client {
	longUidOnce.Do(func() {
		longUid = uid.NewClient(
			8,                // 可拥有 2^8=256 个节点
			10,               // 每个周期生成id上限为 2^10=1024 个
			index,            // 节点编号
			1293811200000,    // 2011-01-01 00:00:00 起始时间戳 控制在18位
			time.Millisecond, // 以1毫秒为一个周期
		)
	})
	return longUid
}

var shortUid *uid.Client
var shortUidOnce sync.Once

func ShortUid(index int64) *uid.Client {
	shortUidOnce.Do(func() {
		shortUid = uid.NewClient(
			4,                   // 可拥有 2^4=16 个节点
			4,                   // 每个周期生成id上限为 2^4=16 个
			index,               // 节点编号
			129381120000,        // 2011-01-01 00:00:00 起始时间戳 控制在14位
			time.Millisecond*10, // 以10毫秒为一个周期
		)
	})
	return shortUid
}

var miniUid *uid.Client
var miniUidOnce sync.Once

func MiniUid(index int64) *uid.Client {
	miniUidOnce.Do(func() {
		miniUid = uid.NewClient(
			2,               // 可拥有 2^2=4 个节点
			6,               // 每个周期生成id上限为 2^6=64 个
			index,           // 节点编号
			12938112,        // 2011-01-01 00:00:00 起始时间戳 控制在10位
			time.Second*100, // 以每100秒为一个周期
		)
	})
	return miniUid
}
