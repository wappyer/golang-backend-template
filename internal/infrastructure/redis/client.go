package redis

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"time"
)

var client *Client

type Client struct {
	client redis.UniversalClient
}

func Initialize(host, port, password string, db int) {
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:       []string{host + ":" + port},
		Password:    password,
		DB:          db,
		IdleTimeout: time.Second * 10,
	})

	// 测试redis连接
	err := c.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Sprintf("redis服务连接失败！请检查：%s:%s", host, port))
	}

	client = &Client{
		client: c,
	}
}

func GetClientIns() *Client {
	return client
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (c *Client) Set(ctx context.Context, key string, val string, expiration time.Duration) error {
	err := c.client.Set(ctx, key, val, expiration).Err()
	return err
}

func (c *Client) Del(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key).Err()
	return err
}
