package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var AppLog *logrus.Logger
var WebLog *logrus.Logger
var SqlLog *logrus.Logger

func Setup() {
	initAppLog()
	initWebLog()
	initSqlLog()
}

//初始化AppLog
func initAppLog() {
	logFileName := "app/" + time.Now().Format("2006-01-02") + ".log"
	AppLog = initLog(logFileName)
}

//初始化WebLog
func initWebLog() {
	logFileName := "web/" + time.Now().Format("2006-01-02") + ".log"
	WebLog = initLog(logFileName)
}

//初始化SqlLog
func initSqlLog() {
	logFileName := "sql/" + time.Now().Format("2006-01-02") + ".log"
	SqlLog = initLog(logFileName)
}

//初始化日志句柄
func initLog(logFileName string) *logrus.Logger{
	log := logrus.New()

	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logPath := "storage/logs/"
	logName := logPath  + logFileName

	var f *os.File
	var err error

	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName,os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		panic("日志文件打开失败")
	}

	log.Out = f
	log.Level = logrus.InfoLevel

	return log
}