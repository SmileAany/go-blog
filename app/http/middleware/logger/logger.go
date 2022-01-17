package logger

import (
	"crm/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		//日志格式
		logger.WebLog.WithFields(logrus.Fields{
			"status": c.Writer.Status(),
			"time" : fmt.Sprintf("%6v",endTime.Sub(startTime)),
			"ip" : c.ClientIP(),
			"method" : c.Request.Method,
			"uri" : c.Request.URL,
		}).Info("send")
	}
}