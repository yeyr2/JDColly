package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"reptile-test-go/config"
)

var db *gorm.DB

func init() {
	ServiceSql()
}

func ServiceSql() {
	var err error
	dsn := config.User + ":" + config.Pass + "@tcp(" + config.DB + ":" + config.SqlPort + ")/jdColly?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("failed to connect database" + fmt.Sprintf("%s", err))
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}
