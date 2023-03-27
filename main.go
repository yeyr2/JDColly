package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"reptile-test-go/api/request"
	_ "reptile-test-go/api/sql"
	"reptile-test-go/setting"
	"time"
)

func init() {
	go setting.Socks5("1080")
	go setting.Socks5("1081")
}

func main() {
	r := gin.Default()

	r.Use(Logger())

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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 记录响应时间和请求路径
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		path := c.Request.URL.Path
		key := c.Request.URL.Query()
		clientIP := c.ClientIP()

		// 打印日志
		c.Writer.WriteHeaderNow()
		logMessage := fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s?%s\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			latencyTime,
			clientIP,
			c.Request.Method,
			path,
			key,
		)

		go writeFile(logMessage)
	}
}

func writeFile(logMessage string) {
	os.MkdirAll("logs", os.ModePerm)

	fileName := fmt.Sprintf("logs%clog-%s", os.PathSeparator, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if !os.IsNotExist(err) {
		write := bufio.NewWriter(file)
		write.WriteString(logMessage)
		write.Flush()
	} else {
		file, err = os.Create(fileName)
		if err != nil {
			log.Println("create file err:", err)
			return
		}
		write := bufio.NewWriter(file)
		write.WriteString(logMessage)
		write.Flush()
	}
}
