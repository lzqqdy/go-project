package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.auth"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.poolsize"),
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}
