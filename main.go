package main

import (
	"github.com/gin-gonic/gin"
	"reptile-test-go/api/request"
	_ "reptile-test-go/api/sql"
	"reptile-test-go/setting"
)

func init() {
	go setting.Socks5("1080")
	go setting.Socks5("1081")
}

func main() {
	r := gin.Default()

	r.Use(setting.Logger())

	initRouter(r)

	err := r.Run(":9090")
	if err != nil {
		panic("run failed.")
	}
}

func initRouter(r *gin.Engine) {
	r.Static("/images", "./images") //图片

	ans := r.Group("/new")

	ans.GET("/colly", request.StartColly)
	ans.GET("/comment", request.GetComment)
	ans.GET("/login", request.Login)
}
