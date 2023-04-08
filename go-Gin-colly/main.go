package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	handler2 "reptile-test-go/handler"
	"reptile-test-go/middleware"
	_ "reptile-test-go/model"
	"time"
)

func init() {
	go middleware.Socks5("1080")
	go middleware.Socks5("1081")
}

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.Logger())

	initRouter(r)

	// 指定服务端口
	addr := ":9090"
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	// 监听并启动服务
	// 利用 goroutine 启动监听
	go func() {
		// srv.ListenAndServe() 监听
		log.Printf("Listening and serving HTTP on %s\n", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	// 反之，则执行后续代码
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用 srv.Shutdown() 完成优雅停止
	// 调用时传递了一个上下文对象，对象中定义了超时时间
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func initRouter(r *gin.Engine) {
	r.Static("/images", "./images") //图片
	r.Static("/wordcloud", "./../pyRpc/images")

	ans := r.Group("/new")

	ans.GET("/colly", handler2.StartColly)
	ans.GET("/comment", handler2.GetComment)
	ans.POST("/login", handler2.Login)
	ans.POST("/register", handler2.Register)
	ans.GET("/search", handler2.SearchInfo)
	ans.GET("/userinfo", handler2.Userinfo)
	ans.GET("/modifyUser", handler2.ModifyUserInformation)
}
