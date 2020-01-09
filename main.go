package main

import (
	"gindemo/config"
	"gindemo/database"
	"gindemo/public"
	"gindemo/router"
)

func main() {
	config.InitConfig()
	database.InitMysql()
	public.InitValidate()
	router := router.InitRouter()
	router.Run()
}
