package bootstrap

import (
	"crm/pkg/model"
	"time"
)

// SetupDatabase 初始化数据库
func SetupDatabase() {
	model.ConnectDatabase()

	sql,err := model.Database.DB()

	if err != nil {
		panic(err.Error())
	}

	sql.SetMaxOpenConns(100)
	sql.SetMaxIdleConns(25)
	sql.SetConnMaxLifetime(5 * time.Minute)
}
