## JDColly

获取京东评论并进行情感分析得出评论分数与词云图片


> 使用了 [中文常用停用词表](https://github.com/goto456/stopwords)  
> 使用docker部署   
> 使用grpc连接go与python项目     
> - go
>   - 使用gin,colly等外部库
>   - 全部依赖在go.mod中
> - python
>   - 使用wordcloud jieba grpcio-tools protobuf snownlp 外部库

> 接口:
>   
