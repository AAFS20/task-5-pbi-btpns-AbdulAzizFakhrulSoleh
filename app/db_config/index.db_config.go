package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_HOST = "127.0.0.1"
var DB_PORT = "3306"
var DB_NAME = "go_db"
var DB_USER = "root"
var DB_PASSWORD = ""

//
// --Membuat file migration
// --migrate create -ext sql -dir database/migrations -seq create_users_table
// --migrate create -ext sql -dir database/migrations create_users_table
// --cara migrasi database
// --migrate -database "mysql://root:@tcp(127.0.0.1:3306)/go_db?param1=true&param2=false" -path database/migrations up

// --migrate -database "mysql://root:@tcp(127.0.0.1:3306)/go_db?param1=true&param2=false" -path database/migrations down
// --Force Dirty Databases
// -- migrate -path database/migrations -database "mysql://root:@tcp(127.0.0.1:3306)/go_db" force 1

func InitDatabaseConfig() {
	driverEnv := os.Getenv("DB_DRIVER")
	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}
	hostEnv := os.Getenv("DB_HOST")
	if hostEnv != "" {
		DB_HOST = hostEnv
	}
	portEnv := os.Getenv("DB_PORT")
	if portEnv != "" {
		DB_PORT = portEnv
	}
	nameEnv := os.Getenv("DB_NAME")
	if nameEnv != "" {
		DB_NAME = nameEnv
	}
	passwordEnv := os.Getenv("DB_PASSWORD")
	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}
}
