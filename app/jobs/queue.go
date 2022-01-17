package jobs

import (
	"crm/app/services"
	"time"
)

var (
	emailChannel = make(chan map[string]string)
)

// Consumption 消费
func Consumption()  {
	for {
		select {
		case v := <- emailChannel:
			services.SendMail([]string{v["email"]},v["subject"],v["body"])
			break
		default:
			time.Sleep(time.Second * 5)
		}
	}
}

// Production 生产
func Production(data map[string]string,name string) {
	switch name {
	case "email":
		emailChannel <- data
		break
	}
}
