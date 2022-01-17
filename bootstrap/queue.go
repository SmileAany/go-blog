package bootstrap

import "crm/app/jobs"

// SetupQueue 初始化队列
func SetupQueue()  {
	jobs.Consumption()
}
