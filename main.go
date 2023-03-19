package main

import (
	"github.com/gin-gonic/gin"
	"reptile-test-go/api"
)

func main() {
	r := gin.Default()

	initRouter(r)

	err := r.Run("www.iyeyr2.top:9090")
	if err != nil {
		panic("run failed.")
	}
}

func initRouter(r *gin.Engine) {
	r.Static("/images", "./images") //图片

	ans := r.Group("/new")

	ans.GET("/colly", api.StartColly)
	ans.GET("/comment", api.GetComment)
	ans.GET("/login", api.Login)
}
