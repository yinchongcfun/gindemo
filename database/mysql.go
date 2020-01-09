package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type Mysql struct {
	UserName string
	Password string
	IpHost   string
	DbName   string
}


type MyLogger struct {
}

func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level  = values[0]
		source = values[1]
	)
	if level == "sql" {
		sql := values[3].(string)
		log.Println(sql, level, source)
	} else {
		log.Println(values)
	}

}

var GormPool *gorm.DB

func InitMysql() {
	var err error
	GormPool, err = gorm.Open("mysql", "blog:123456@tcp(118.24.93.185:3306)/blog?parseTime=true")

	//logger:=&MyLogger{}
	GormPool.LogMode(true)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if GormPool.Error != nil {
		fmt.Printf("database error %v", GormPool.Error)
	}
}
