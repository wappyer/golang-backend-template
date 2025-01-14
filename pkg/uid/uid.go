package uid

import (
	"strconv"
	"sync"
	"time"
)

/*
 * 雪花算法的唯一id生成器 适用于分布式的系统id生成
 * 原理即维护一个二进制数，高位存放时间周期，中位存放workerId节点id，低位为number序号
 * 生成id时会加互斥锁
 * 同一节点同一时间周期内最多可生成id数为序号位能存的最大数，超过的话会等待至下一周期
 */

type Client struct {
	mu sync.Mutex // 添加互斥锁 确保并发安全

	WorkerId   int64 // 该节点的ID
	WorkerBits uint8 // 节点数，可拥有 2^n-1 个节点
	workerMax  int64 // 节点ID的最大值，防止溢出

	NumberBits uint8 // 每个时间周期(durationType)可生成 2^n-1 个唯一ID
	numberMax  int64 // 序号ID的最大值，防止溢出
	number     int64 // 当前毫秒已经生成的id序列号(从0开始累加)

	TimeEpoch    int64         // 开始时间戳
	TimeCycle    int64         // 当前时间周期
	TimeDuration time.Duration // 时间周期（如配置time.Millisecond,则以每毫秒一个周期限制number的数量）
}

func NewClient(workerBits, numberBits uint8, workerId, timeEpoch int64, timeDuration time.Duration) *Client {
	client := &Client{
		WorkerId:     workerId,
		WorkerBits:   workerBits,
		NumberBits:   numberBits,
		TimeEpoch:    timeEpoch,
		TimeDuration: timeDuration,
	}
	client.numberMax = -1 ^ (-1 << client.NumberBits) // 序号ID的最大值，防止溢出
	client.workerMax = -1 ^ (-1 << client.WorkerBits) // 节点ID的最大值，防止溢出
	return client
}

func (c *Client) GetStr() string {
	return strconv.Itoa(int(c.GetInt64()))
}

func (c *Client) GetInt64() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 当前时间周期
	now := time.Now().UnixNano() / int64(c.TimeDuration)

	if c.TimeCycle == now {
		c.number++
		// 如果当前工作节点在一个时间周期内生成的ID已经超过上限 需要等到下个时间周期再继续生成
		if c.number > c.numberMax {
			for now <= c.TimeCycle {
				now = time.Now().UnixNano() / int64(c.TimeDuration)
			}
		}
	} else {
		// 重置序号number与时间周期timeCycle
		c.number = 0
		c.TimeCycle = now
	}

	workerIdOffset := c.NumberBits                 // 节点id向左偏移位（即为序号number的位数）
	timeCycleOffset := c.WorkerBits + c.NumberBits // 周期数向左偏移位
	timeCycle := now - c.TimeEpoch                 // 距起始周期数

	ID := timeCycle<<timeCycleOffset | c.WorkerId<<workerIdOffset | c.number
	return ID
}
