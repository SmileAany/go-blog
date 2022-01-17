package config

import "crm/pkg/config"

func init() {
	config.Add("database", config.StrMap{
		"host" : config.Env("DB_HOST", "127.0.0.1"),
		"database" : config.Env("DB_DATABASE", "database"),
		"username" : config.Env("DB_USERNAME", "username"),
		"password" : config.Env("DB_PASSWORD", "password"),
		"charset"  : "utf8",
		"parseTime" : "true",
		"loc" : "Local",
  	})
}
