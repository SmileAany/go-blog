package config

import "crm/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		"port": config.Env("APP_PORT", "8080"),
	})
}
