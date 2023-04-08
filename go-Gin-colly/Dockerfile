FROM golang:1.19

# 设置工作目录
WORKDIR /JDColly

# 复制 Go 程序代码到镜像中
COPY . .

RUN  go mod tidy && \
     go mod download && \
     go build main.go

# 暴露端口
EXPOSE 9090 80 50051 50052

# 启动应用程序
CMD ["./main"]