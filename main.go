package main

import (
	"log"

	"github.com/aafs20/rakamin_golang/app"
	"github.com/aafs20/rakamin_golang/app/app_config"
	"github.com/aafs20/rakamin_golang/database"
	"github.com/aafs20/rakamin_golang/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() //LOAD .ENV FILE
	if err != nil {
		log.Println("Error Loading .env file")
	}

	app.InitConfig()            //INIT CONFIG
	database.ConnectDatabases() //DATABASE CONNECTION
	app := gin.Default()        //INIT GIN ENGINE

	router.InitRoute(app) //INIT ROUTE

	app.Run(app_config.PORT) //RUN APP
}
