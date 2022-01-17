package config

import "crm/pkg/config"

func init() {
	config.Add("email", config.StrMap{
		"from" : config.Env("EMAIL_FROM", "email"),
		"host" : config.Env("EMAIL_HOST","127.0.01"),
		"port" : config.Env("EMAIL_PORT","25"),
		"username" : config.Env("EMAIL_USERNAME","username"),
		"password" : config.Env("EMAIL_PASSWORD","password"),
		"ssl" : config.Env("EMAIL_SSL",false),
	})
}
