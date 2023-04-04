package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"reptile-test-go/api/goGRPC/wordsCloud"
)

func main() {
	// 创建gRPC连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// 创建Greeter客户端
	client := wordsCloud.NewGreeterClient(conn)
	request := &wordsCloud.RpcComment{Content: []string{"早上好", "下午好下午打", "晚上好"}}
	response, err := client.WordCloudAnalysis(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	//analyze.AnalyzeWord = response.WordsCloud
	// 打印响应的message属性
	fmt.Println(response.GetWordsCloud())

}
