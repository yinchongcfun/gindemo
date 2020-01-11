package main

import (
	"gindemo/config"
	"gindemo/database"
	"gindemo/public"
	"gindemo/router"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	config.InitConfig()
	database.InitMysql()
	database.InitRedis()
	public.InitValidate()
	router := router.InitRouter()
	router.Run()
}
