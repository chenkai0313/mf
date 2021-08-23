package app

import (
	"fmt"

	"github.com/go-redis/redis"

	"api/config"
)

var RedisDB *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "" + config.GetRedis().HostName + ":" + config.GetRedis().Port + "",
		Password: "" + config.GetRedis().Password + "",
		DB:       config.GetRedis().DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis connect error= %v", err))
	}
	RedisDB = client
}
