# 使用官方 Python 镜像作为基础镜像
FROM python:3.9

# 设置工作目录
WORKDIR /pyWordCloud

# 复制 Python 程序代码到镜像中
COPY . .

RUN  pip3 install snownlp

# 暴露端口
EXPOSE 50052 9090 80

# 指定 Python 程序运行时的命令
CMD [ "python", "service.py" ]
