package main

import (
	"crm/bootstrap"
	c "crm/pkg/config"
	"crm/routers"
)

func init() {
	//初始化配置文件
	bootstrap.SetupConfig()

	//初始化日志
	bootstrap.SetupLogger()

	//初始化数据库
	bootstrap.SetupDatabase()

	//初始化redis
	bootstrap.SetupRedis()

	//初始化队列
	go bootstrap.SetupQueue()
}

func main() {
	route := routers.SetupRouter()

	route.Run(":" + c.GetString("app.port"))
}
