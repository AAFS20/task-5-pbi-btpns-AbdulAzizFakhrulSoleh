package app

import (
	"github.com/aafs20/rakamin_golang/app/app_config"
	"github.com/aafs20/rakamin_golang/app/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}
