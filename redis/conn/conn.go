package conn

import (
	"github.com/go-redis/redis"
)

func NewRedisClient(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := client.Ping().Err(); err != nil {
		panic("Connect Redis failed.")
	}
	return client
}
