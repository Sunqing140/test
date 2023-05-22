#从golang1.14构建
FROM golang:1.14 AS app 
#开启模块支持
ENV GO111MODULE=on  
#换源
ENV GOPROXY https://goproxy.cn,direct
#切换目录
WORKDIR /root/togettoyou
#复制服务器资源
COPY server/. .
#构建app
RUN CGO_ENABLED=0 go build -v -o "app" .

#从node构建
FROM node:12 AS web
#切换目录
WORKDIR /root/togettoyou
#复制需要安装的包信息
COPY web/package*.json ./
#安装包
RUN yarn config set registry https://registry.npmmirror.com/ && yarn install
#复制前端资源
COPY web/. .
#启动前端web
RUN yarn run build:prod


#空镜像
FROM scratch
#切换目录
WORKDIR /root/togettoyou/
#复制一些配置信息
COPY --from=app /root/togettoyou/app ./
COPY --from=app /root/togettoyou/config.yaml ./
COPY --from=web /root/togettoyou/dist/ ./dist/
#docker镜像运行的命令，app就是这个项目
ENTRYPOINT ["./app"]
