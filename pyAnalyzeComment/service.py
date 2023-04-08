from concurrent import futures
import grpc

import service.analyze_pb2_grpc as analyze_pb2_grpc
import analysisWords.vaderAnalysis as analyze

class AnalyzeServicer(analyze_pb2_grpc.GreeterServicer):
    def AnalysisCommentsByNLP(self, request, context):
        # 在这里编写处理业务逻辑的代码
        response = analyze.analysis_comment(request.Content)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    analyze_pb2_grpc.add_GreeterServicer_to_server(AnalyzeServicer(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
