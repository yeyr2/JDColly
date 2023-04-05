# 使用官方 Python 镜像作为基础镜像
FROM python:3.8

# 设置工作目录
WORKDIR /pyRpc

# 复制 Python 程序代码到镜像中
COPY . .

# 安装 Python 依赖模块
RUN pip3 install -r requirements.txt

# 暴露端口
EXPOSE 50051

# 指定 Python 程序运行时的命令
CMD [ "python", "service.py" ]
