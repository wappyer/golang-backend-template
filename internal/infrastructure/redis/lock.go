package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// SetLock 加分布式锁
func (c *Client) SetLock(ctx context.Context, key string, val string) (bool, error) {
	args := redis.SetArgs{
		Mode: "NX",
		TTL:  time.Second * 10,
	}
	_, err := c.client.SetArgs(ctx, key, val, args).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// DelLock 删分布式锁
func (c *Client) DelLock(ctx context.Context, key string, lockId string) (interface{}, error) {
	// redis 执行 lua 脚本删除
	script := `local id = redis.call('get', KEYS[1]); 
if id == ARGV[1]
then 
	return redis.call('del', KEYS[1]);
end
return 2`
	ret, err := c.client.Eval(ctx, script, []string{key}, lockId).Result()
	return ret, err
}
