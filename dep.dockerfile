## 各服务的依赖镜像, 用来缩短构建具体服务时下载依赖的时间

FROM golang:latest
WORKDIR ${GOPATH}/src/github.com/zsy-cn/bms
RUN rm -rf ./*
COPY . .

RUN set -x && go get -d -v ./loraclient
RUN set -x && go get -d -v ./server