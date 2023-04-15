# 使用官方 Node.js 映像作为基础映像
FROM node:14

# 设置工作目录
WORKDIR /web

# 复制 package.json 和 package-lock.json 文件到镜像中
COPY package*.json ./

# 安装依赖项
RUN npm install -g cnpm --registry=http://registry.npm.taobao.org && \
    npm install -g @angular/cli && \
    cnpm install -g @vue/cli && \
    npm install

# 将应用程序代码复制到镜像中
COPY . .

# 暴露容器端口
EXPOSE 8080

# 启动应用程序
CMD ["export","DANGEROUSLY_DISABLE_HOST_CHECK=true"]
CMD [ "npm", "start" ]
