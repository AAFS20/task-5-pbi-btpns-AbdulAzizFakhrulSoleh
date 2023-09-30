package database

import (
	"log"

	"github.com/aafs20/rakamin_golang/app/db_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabases() {
	var errConnection error
	if db_config.DB_DRIVER == "mysql" {
		dsnMysql := "root:@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True&loc=Local"
		DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})

	}
	if errConnection != nil {
		panic("Can't connect to database")
	}

	log.Println("Connected to database")
}
