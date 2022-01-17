package redis

import (
	"crm/pkg/config"
	"github.com/go-redis/redis"
)

var Client * redis.Client

func ConnectRedis() * redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr: config.GetString("redis.host") + ":" + config.GetString("redis.port"),
		Password : config.GetString("redis.password"),
		DB : config.GetInt("redis.database"),
	})

	return Client
}