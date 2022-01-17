package config

import "crm/pkg/config"

func init() {
	config.Add("redis", config.StrMap{
		"port" : config.Env("REDIS_PORT", "6379"),
		"host" : config.Env("REDIS_HOST","127.0.0.1"),
		"password" : config.Env("REDIS_PASSWORD","password"),
		"database" : config.Env("REDIS_DATABASE","0"),
	})
}
