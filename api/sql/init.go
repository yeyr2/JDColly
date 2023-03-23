package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"reptile-test-go/setting"
)

var db *gorm.DB

func init() {
	ServiceSql()
}

func ServiceSql() {
	var err error
	dsn := setting.User + ":" + setting.Pass + "@tcp(127.0.0.1:3306)/jdColly?charaset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("failed to connect database" + fmt.Sprintf("%s", err))
	}
}
