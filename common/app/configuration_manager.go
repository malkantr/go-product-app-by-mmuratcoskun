package app

import "product-app/common/postgresql"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: postgreSqlConfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "5432",
		UserName:              "postgres",
		Password:              "Password.1",
		DbName:                "go_productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
