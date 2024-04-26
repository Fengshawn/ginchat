FROM golang:1.22.0

ARG proxy=https://goproxy.cn
ENV proxy=${proxy}
RUN echo ${proxy}


# 设置工作目录
WORKDIR /app

ENV GOPROXY ${proxy}

# 拷贝go.mod和go.sum以下载依赖
COPY go.mod .
COPY go.sum .

# 下载依赖
RUN go mod download

# 拷贝应用程序源代码到容器中
COPY . .

#RUN apt-get update && apt-get install -y net-tools && apt-get install -y vim

# 构建应用程序
#RUN go build -o main .

#暴露端口
EXPOSE 8081

# 定义容器启动时执行的命令
# CMD ["tail", "-f", "/dev/null"] #死循环调试容器环境用

CMD ["go", "run", "main.go"]
# CMD ["sh", "-c", "sleep 1 && go run main.go"]
