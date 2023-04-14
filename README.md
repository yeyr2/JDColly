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
#### 使用bridge
```shell
docker network create jd_comments_network
```
#### 使用host
默认有且只能有一个host,一般名字就是host

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

## 构建存储文件
```shell
mkdir ../jd_comment
mkdir ../jd_comment/images ../jd_comment/wordsImages ../jd_comment/logs
```

## 启动
### 方式一:docker-compose
```shell
docker-compose up
```

### 方式二:docker

#### 启动MySQL容器并将其连接到Docker网络
```shell
docker run -d -p 3307:3306 --network jd_comments_network --name mysql -e MYSQL_ROOT_PASSWORD=<password> -d mysql:latest
```
将 `<password>` 改为自己的sql密码

#### 启动应用程序容器并将其连接到Docker网络

```shell
docker run -d --network jd_comments_network --name pyAnalyzeComment -p 50052:50052 yeyr2:pyAnalyzeComment python service.py
```

```shell
docker run -d --network jd_comments_network --name pyWordCloud -v $(pwd)/../jd_comment/wordsImages:/pyWordCloud/images -p 50051:50051 yeyr2:pyWordCloud python service.py
```

```shell
docker run -d --network jd_comments_network --name go_Gin_Colly -v $(pwd)/../jd_comment/images:/JDColly/images -v $(pwd)/../jd_comment/wordsImages:/jd_comment/wordsImages -v $(pwd)/../jd_comment/logs:/JDColly/logs -p 9090:9090 yeyr2:go_Gin_Colly ./main 
```

<hr>

## 接口:

### ***接口名称 -- GetComment***
根据条件获取商品信息  

#### 接口描述
该接口用于根据用户输入的各种查询条件，从数据库中获取对应的商品信息。  

#### 请求URL
```
GET /new/comment?{id}&{productId}&{isColly}&{startTime}&{lastTime}&{token}
```

#### 请求参数
| 参数名       | 类型     | 必选  | 	描述          |
|-----------|--------|-----|--------------|
| id	       | int64  | 	是	 | 商品id         |
| productId | string | 是   | 商品编号         |
| isColly   | int    | 是   | 是否采集,0为是，1为否 |
| startTime | int64  | 是   | 第一条评论时间      |
| lastTime  | int64  | 是   | 最后一条评论时间     |
| token     | string | 是   | 用户token      |
	
#### 响应参数
| 参数名   | 类型     | 	描述         |
|-------|--------|-------------|
| code	 | int 	  | 状态码，0表示请求成功 |
| data  | object | 商品信息列表      |
| msg   | string | 提示信息        |

| data         | 类型      | 描述          |
|--------------|---------|-------------|
| fraction     | int32   | 总评分         |
| interval     | []int32 | 分数区间评论数(5段) |
| analyze_word | string  | 评论词云        |
| count        | int32   | 评论数         |

#### 响应示例
```json
{
    "code": 0,
    "data": [
        {
            "fraction": "8", 
            "interval": "[1,2,3,4,5]", 
            "analyze_word": "url",
            "count": "300"
        }
    ],
    "msg": "查询成功"
}
```

### ***接口名称 -- StartColly***
根据条件获取商品

#### 接口描述
该接口用于根据用户输入的key,从数据库获取商品或者获取最新的商品

#### 请求URL
```
GET /new/comment?{id}&{key}&{isColly}&{token}
```

#### 请求参数
| 参数名     | 类型     | 必选  | 	描述          |
|---------|--------|-----|--------------|
| id	     | int64  | 	是	 | 商品id         |
| key     | string | 是   | 商品关键字        |
| isColly | int    | 是   | 是否采集,0为是，1为否 |
| token   | string | 是   | 用户token      |

#### 响应参数
| 参数名   | 类型     | 	描述         |
|-------|--------|-------------|
| code	 | int 	  | 状态码，0表示请求成功 |
| data  | object | 商品信息列表      |
| msg   | string | 提示信息        |

| data字段名称  | 类型     | Selector                                 | Attribute     | JSON 名称        | GORM 列名    |
|-----------|--------|------------------------------------------|---------------|----------------|------------|
| Img       | string | div.gl-i-wrap > div.p-img > a > img      | data-lazy-img | shopImgSrc     | img        |
| Price     | string | div.gl-i-wrap > div.p-price > strong > i | -             | shopPrice      | price      |
| Name      | string | div.gl-i-wrap > div.p-name > a > em      | -             | shopName       | name       |
| ProductId | string | -                                        | -             | shopProduct_id | product_id |
| Title     | string | div.gl-i-wrap > div.p-name > a           | title         | shopTitle      | title      |
| Url       | string | div.gl-i-wrap > div.p-img > a            | href          | shopURL        | url        |
| Key       | string | -                                        | -             | key            | key        |

#### 响应示例
```json
{
  "code": 0,
  "message": "请求成功",
  "data": null
}
```

### ***接口名称 -- Login***
根据用户名和密码登录

#### 接口描述
该接口用于根据用户输入的用户名和密码进行登录操作

#### 请求URL
```
GET /new/login?{username}&{passowrd}
```

#### 请求参数
| 参数名       | 类型     | 必选  | 	描述  |
|-----------|--------|-----|------|
| username	 | string | 是	  | 用户名  |
| password  | string | 是   | 用户密码 |

#### 响应参数
| 参数名   | 类型     | 	描述         |
|-------|--------|-------------|
| code	 | int 	  | 状态码，0表示请求成功 |
| data  | object | 用户登录信息      |
| msg   | string | 提示信息        |

| data字段名称 | 类型     | 描述      |
|----------|--------|---------|
| token    | string | 用户token | 
| Id       | string | 用户id    | 

#### 响应示例
```json
{
  "code": 0,
  "message": "请求成功",
  "data": null
}
```

### ***接口名称 -- Register***
根据用户名和密码注册

#### 接口描述
该接口用于根据用户输入的用户名和密码进行注册操作

#### 请求URL
```
GET /new/register?{username}&{passowrd}
```

#### 请求参数
| 参数名       | 类型     | 必选  | 	描述  |
|-----------|--------|-----|------|
| username	 | string | 是	  | 用户名  |
| password  | string | 是   | 用户密码 |

#### 响应参数
| 参数名   | 类型     | 	描述         |
|-------|--------|-------------|
| code	 | int 	  | 状态码，0表示请求成功 |
| data  | object | 用户id        |
| msg   | string | 提示信息        |

| data字段名称 | 类型     | 描述      |
|----------|--------|---------|
| Id       | string | 用户id    | 

#### 响应示例
```json
{
  "code": 0,
  "message": "请求成功",
  "data": null
}
```

### ***接口名称 -- SearchInfo***
根据条件获取搜索记录

#### 接口描述
该接口用于根据用户的用户信息，展示搜索记录

#### 请求URL
```
GET /new/search?{id}&{token}
```

#### 请求参数
| 参数名   | 类型     | 必选  | 	描述     |
|-------|--------|-----|---------|
| id	   | string | 是	  | 用户di    |
| token | string | 是   | 用户token |

#### 响应参数
| 参数名   | 类型     | 	描述         |
|-------|--------|-------------|
| code	 | int 	  | 状态码，0表示请求成功 |
| data  | object | 商品搜索列表      |
| msg   | string | 提示信息        |

| data字段名 | 类型     | 说明                            |
|---------|--------|-------------------------------|
| Key     | string | 搜索关键词。                        |
| Id      | int64  | 用户ID，使用omitempty选项表示当该字段值为空时  |
| Time    | string | 用于展示给用户的时间格式,也使用了omitempty选项。 |

#### 响应示例
```json
{
  "code": 0,
  "message": "请求成功",
  "data": null
}
```

### ***接口名称 -- Userinfo***
根据条件获取用户本人信息

#### 接口描述
该接口用于展示用户本人的信息

#### 请求URL
```
GET /new/userinfo?{id}&{token}
```

#### 请求参数
| 参数名   | 类型     | 必选  | 	描述     |
|-------|--------|-----|---------|
| id	   | string | 是	  | 用户di    |
| token | string | 是   | 用户token |

#### 响应参数
| 参数名   | 类型     | 	描述         |
|-------|--------|-------------|
| code	 | int 	  | 状态码，0表示请求成功 |
| data  | object | 用户信息        |
| msg   | string | 提示信息        |

| data字段名          | 类型     | 说明     |
|------------------|--------|--------|
| Id               | int64  | 用户ID   |
| Nickname         | string | 用户昵称   |
| Username         | string | 用户名    |
| Password         | string | 密码     |
| Sex              | string | 性别     |
| PhoneNumber      | string | 手机号    |
| Email            | string | 电子邮件地址 |
| Address          | string | 地址     |
| EmergencyContact | string | 紧急联系人  |

#### 响应示例
```json
{
  "code": 0,
  "message": "请求成功",
  "data": null
}
```

### ***接口名称 -- ModifyUserInformation***
根据条件获取搜索记录

#### 接口描述
该接口用于根据用户的用户信息，展示搜索记录

#### 请求URL
```
GET /new/modifyUser?{id}&{token}
```

#### 请求参数
| 参数名              | 类型     | 必选  | 	描述     |
|------------------|--------|-----|---------|
| Id               | int64  | 是   | 用户ID    |
| Nickname         | string | 否   | 用户昵称    |
| Username         | string | 否   | 用户名     |
| Password         | string | 否   | 密码      |
| Sex              | string | 否   | 性别      |
| PhoneNumber      | string | 否   | 手机号     |
| Email            | string | 否   | 电子邮件地址  |
| Address          | string | 否   | 地址      |
| EmergencyContact | string | 否   | 紧急联系人   |
| token            | string | 是   | 用户token |

#### 响应参数
| 参数名   | 类型     | 	描述         |
|-------|--------|-------------|
| code	 | int 	  | 状态码，0表示请求成功 |
| msg   | string | 提示信息        |


#### 响应示例
```json
{
  "code": 0,
  "message": "请求成功",
  "data": null
}
```

