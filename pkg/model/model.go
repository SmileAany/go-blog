package model

import (
	c "crm/pkg/config"
	log "crm/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	Database *gorm.DB
)

// ConnectDatabase ConnectDB 数据库初始化
func ConnectDatabase() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v&parseTime=%v&loc=%v",
		c.GetString("database.username"),
		c.GetString("database.password"),
		c.GetString("database.host"),
		c.GetString("database.database"),
		c.GetString("database.charset"),
		c.GetString("database.parseTime"),
		c.GetString("database.loc"),
	)

	config := mysql.New(mysql.Config{
		DSN: dsn,
	})

	//绑定sql日志
	gormLogger := logger.New(log.SqlLog,logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful: true,
	})

	database, err := gorm.Open(config, &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil || database == nil {
		panic("数据库连接失败")
	}

	Database = database
}
