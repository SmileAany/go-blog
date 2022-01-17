package services

import (
	"crm/pkg/config"
	"crm/pkg/logger"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"net/smtp"
)

type Email struct {
	email.Email
	Host     string
	Port     string
	UserName string
	Password string
	From     string
}

func SendMail(to []string, subject, body string) error {
	parameter := &Email{
		Host:     config.GetString("email.host"),
		Port:     config.GetString("email.port"),
		UserName: config.GetString("email.username"),
		Password: config.GetString("email.password"),
		From:     config.GetString("email.from"),
	}

	auth := smtp.PlainAuth("",parameter.UserName,parameter.Password,parameter.Host)

	e := &email.Email{
		From:    fmt.Sprintf(parameter.From+"<%s>", parameter.UserName),
		To:      to,
		Subject: subject,
		HTML:    []byte(body),
	}

	err := e.Send(parameter.Host+parameter.Port, auth)

	if err != nil {
		//日志格式
		logger.AppLog.WithFields(logrus.Fields{
			"error" : err.Error(),
		}).Info("email")

		return err
	}

	return nil
}