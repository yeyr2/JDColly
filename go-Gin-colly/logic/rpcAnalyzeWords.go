package logic

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"reptile-test-go/cmd"
	"reptile-test-go/goGRPC/AnalysisByNLP"
)

var analyzeWordsClient AnalysisByNLP.GreeterClient

func init() {
	// 创建gRPC连接
	conn, err := grpc.Dial("pyAnalyzeComment:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	// 创建Greeter客户端
	analyzeWordsClient = AnalysisByNLP.NewGreeterClient(conn)
}

func AnalysisByNLPRpc(comment *[]cmd.Comments) (int32, []int32) {
	request := cmdFormNLPRPC(comment)
	response, err := analyzeWordsClient.AnalysisCommentsByNLP(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	return response.Fraction, response.Interval
}

func cmdFormNLPRPC(comment *[]cmd.Comments) *AnalysisByNLP.RpcComment {
	ans := new(AnalysisByNLP.RpcComment)
	ans.Content = make([]string, 0, len(*comment))
	for _, x := range *comment {
		ans.Content = append(ans.Content, x.Content)
	}
	return ans
}
