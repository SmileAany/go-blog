package bootstrap

import "crm/pkg/redis"

// SetupRedis 初始化队列
func SetupRedis()  {
	redis.ConnectRedis()
}