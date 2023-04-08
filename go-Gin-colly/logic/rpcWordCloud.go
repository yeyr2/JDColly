package logic

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"reptile-test-go/cmd"
	WordsCloud2 "reptile-test-go/goGRPC/WordsCloud"
)

var wordCloudClient WordsCloud2.GreeterClient

func init() {
	// 创建gRPC连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	// 创建Greeter客户端
	wordCloudClient = WordsCloud2.NewGreeterClient(conn)
}

func wordCloudRpc(comment *[]cmd.Comments, id string) string {
	request := cmdFormRPC(comment)
	request.ProductId = id
	response, err := wordCloudClient.WordCloudAnalysis(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	return response.WordsCloud
}

func cmdFormRPC(comment *[]cmd.Comments) *WordsCloud2.RpcComment {
	ans := new(WordsCloud2.RpcComment)
	ans.Content = make([]string, 0, len(*comment))
	for _, x := range *comment {
		ans.Content = append(ans.Content, x.Content)
	}
	return ans
}
