package main

import (
	"gindemo/config"
	"gindemo/database"
	_ "gindemo/docs"
	hook "gindemo/log"
	"gindemo/public"
	"gindemo/router"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func initLog() {
	uuids := uuid.NewV1()
	log.AddHook(hook.NewTraceIdHook(uuids.String() + " "))
	timeStr := time.Now().Format("2006-01-02")
	file, _ := os.OpenFile("log/"+timeStr+".log", os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	//设置最低loglevel
	log.SetLevel(log.InfoLevel)
}

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目
// @contact.name
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
func main() {
	initLog()
	config.InitConfig()
	database.InitMysql()
	database.InitRedis()
	public.InitValidate()
	router := router.InitRouter()
	router.Run()
}
