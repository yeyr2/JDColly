package api

import (
	"google.golang.org/grpc"
	"log"
	"reptile-test-go/api/goGRPC/wordsCloud"
)

var client wordsCloud.GreeterClient

func init() {
	// 创建gRPC连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// 创建Greeter客户端
	client = wordsCloud.NewGreeterClient(conn)
}
