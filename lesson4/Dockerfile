FROM golang:1.16-alpine
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"
# docker中的工作目录
WORKDIR /myhttpserver
# 将当前目录同步到docker工作目录下
ADD . ./
# 将代码编译成二进制可执行文件  可执行文件名为 myserver
RUN go build -o myserver main.go
EXPOSE 8080
ENTRYPOINT  ["./myserver"]