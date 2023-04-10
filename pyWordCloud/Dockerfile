# 使用官方 Python 镜像作为基础镜像
FROM python:3.9

# 设置工作目录
WORKDIR /pyWordCloud

# 复制 Python 程序代码到镜像中
COPY . .

# 设置容器网络为主机网络

#RUN apt-get update && apt-get install -y net-tools iputils-ping curl && \
#    python -m pip install --upgrade pip && \
#    pip install -i https://mirrors.aliyun.com/pypi/simple/ flask && \
#    apt-get clean && \
#    rm -rf /var/lib/apt/lists/*

# 安装 Python 依赖模块
#RUN pip3 install --no-cache-dir -r requirements.txt && \
#    apt-get update && \
#    apt-get install -y libproj-dev proj-data proj-bin libgeos-dev && \
#    apt-get clean && \
#    rm -rf /var/lib/apt/lists/* && \
RUN  python -m pip install --upgrade pip && \
pip3 install wordcloud jieba grpcio-tools protobuf

# RUN pip3 install -r requirements.txt

# 暴露端口
EXPOSE 50051

# 指定 Python 程序运行时的命令
CMD [ "python", "service.py" ]
