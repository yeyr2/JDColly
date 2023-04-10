## JDColly

获取京东评论并进行情感分析得出评论分数与词云图片(ing...)


> 使用了 [中文常用停用词表](https://github.com/goto456/stopwords)  
> 使用docker部署   
> 使用grpc连接go与python项目     
>   - go
>     - 使用gin,colly等外部库
>     - 全部依赖在go.mod中
>   - python
>     - 使用wordcloud jieba grpcio-tools protobuf snownlp 外部库

# 构建docker镜像

```shell
docker build -t yeyr2:go_Gin_Colly ./go-Gin-colly
docker build -t yeyr2:pyAnalyzeComment ./pyAnalyzeComment
docker build -t yeyr2:pyWordCloud ./pyWordCloud
```
s
# 启动

```shell
mkdir ../jd_comment/images
mkdir ../jd_comment/wordsImages
mkdir ../jd_comment/logs
docker run -v ../jd_comment/images:./images -v ../jd_comment/wordsImages:./logs -p 9090:9090 -p 50051:50051 -p 50052:50052 yeyr2:go_Gin_Colly ./main 
docker run -v ../jd_comment/wordsImages:./images -p 50051:50051 yeyr2:pyWordCloud python service.py
docker run -p 50052:50052 yeyr2:pyAnalyzeComment python service.py
```
    

# 接口:
    
