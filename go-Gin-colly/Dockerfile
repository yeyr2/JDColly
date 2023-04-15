FROM golang:1.19

# 设置工作目录
WORKDIR /JDColly

# 复制 Go 程序代码到镜像中
COPY . .

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone && \
    go mod tidy && \
    go mod download && \
    go build main.go && \
    rm -f /etc/localtime && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 暴露端口
EXPOSE 9090

# 启动应用程序
CMD ["./main"]