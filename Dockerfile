FROM golang:alpine
MAINTAINER lzqqdy <lzqqdy@qq.com>
ENV GOPROXY=https://goproxy.cn,direct
# 移动到工作目录：/gowork
WORKDIR /gowork

# 将代码复制到容器中
COPY . .

EXPOSE 8080
