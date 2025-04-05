package core

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

func InitRedis(addr, pwd string, db int) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
		PoolSize: 100, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
		return
	}
	return
}
