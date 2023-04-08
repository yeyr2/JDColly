from concurrent import futures

import grpc

import service.wordCloud_pb2 as wordCloud_pb2
import service.wordCloud_pb2_grpc as wordCloud_pb2_grpc
from WordsCloud import wordCloud


class GreeterServicer(wordCloud_pb2_grpc.GreeterServicer):
    def WordCloudAnalysis(self, request, context):
        # request.Content 是一个字符串列表
        # 并返回一个 wordCloud_pb2.rpcWordCloud 对象
        string = str(wordCloud.split_text(text=request.Content))
        path = wordCloud.word_cloud(string, productId=request.productId)
        return wordCloud_pb2.rpcWordCloud(wordsCloud=path)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    wordCloud_pb2_grpc.add_GreeterServicer_to_server(GreeterServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
