package api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"reptile-test-go/api/cmd"
	"reptile-test-go/api/goGRPC/wordsCloud"
)

var client wordsCloud.GreeterClient

func init() {
	// 创建gRPC连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	// 创建Greeter客户端
	client = wordsCloud.NewGreeterClient(conn)
}

func wordCloudRpc(comment *[]cmd.Comments) string {
	request := cmdFormRPC(comment)
	response, err := client.WordCloudAnalysis(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	return response.WordsCloud
}

func cmdFormRPC(comment *[]cmd.Comments) *wordsCloud.RpcComment {
	ans := new(wordsCloud.RpcComment)
	ans.Content = make([]string, 0, len(*comment))
	for _, x := range *comment {
		ans.Content = append(ans.Content, x.Content)
	}
	return ans
}
