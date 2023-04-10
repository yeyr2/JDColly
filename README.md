# JDColly

获取京东评论并进行情感分析得出评论分数与词云图片(ing...)


> 使用了 [中文常用停用词表](https://github.com/goto456/stopwords)  
> 使用docker部署   
> 使用grpc连接go与python项目     
>   - go
>     - 使用gin,colly等外部库
>     - 全部依赖在go.mod中
>   - python
>     - 使用wordcloud jieba grpcio-tools protobuf snownlp 外部库

> 使用前请修改`go-Gin-colly/config/setting.go`文件

## 构建docker镜像

### 创建一个Docker网络     
```shell
docker network create jd_comments_network
```

### 构建镜像   
```shell
cd ./go-Gin-colly
docker build -t yeyr2:go_Gin_Colly .
cd ..
```

```shell
cd ./pyAnalyzeComment
docker build -t yeyr2:pyAnalyzeComment .
cd ..
```

```shell
cd ./pyWordCloud
docker build -t yeyr2:pyWordCloud .
cd ..
```

如果下载时超时,可以使用本机的网络
```shell
cd ./go-Gin-colly
docker build --network="host" -t yeyr2:go_Gin_Colly .
cd ..
```

```shell
cd ./pyAnalyzeComment
docker build --network="host" -t yeyr2:pyAnalyzeComment .
cd ..
```

```shell
cd ./pyWordCloud
docker build --network="host" -t yeyr2:pyWordCloud .
cd ..
```

## 启动

### 启动MySQL容器并将其连接到Docker网络  
```shell
docker run -d -p 9006:3306 --network jd_comments_network --name mysql -e MYSQL_ROOT_PASSWORD=<password> -d mysql:latest
```
将 `<password>` 改为自己的sql密码

### 构建存储文件
```shell
mkdir ../jd_comment
mkdir ../jd_comment/images ../jd_comment/wordsImages ../jd_comment/logs
```

### 启动应用程序容器并将其连接到Docker网络
```shell
docker run -d --network jd_comments_network -e DB_HOST=host.docker.internal -e DB_PORT=3306 -e DB_USER=root -e DB_PASSWORD=<password> -v $(pwd)/../jd_comment/images:/jd_comment/images -v $(pwd)/../jd_comment/wordsImages:/jd_comment/logs -p 9090:9090 yeyr2:go_Gin_Colly ./main 
```
`<password>`为你的mysql密码
```shell
docker run -d --network jd_comments_network -p 50052:50052 yeyr2:pyAnalyzeComment python service.py
```
```shell
docker run -d --network jd_comments_network -v $(pwd)/../jd_comment/wordsImages:/pyWordCloud/images -p 50051:50051 yeyr2:pyWordCloud python service.py
```



## 接口:
    
