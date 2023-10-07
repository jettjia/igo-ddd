# Relying on Alpine go version 1.20
FROM --platform=$TARGETPLATFORM golang:1.20-alpine AS builder
MAINTAINER jettjia <jettjia@qq.com>

# Pull private repo go-lib
RUN apk add --no-cache git
RUN echo "machine github.com login jettjia password your_pwd" >>  ~/.netrc

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY ../../go.mod .
COPY ../../go.sum .
RUN go mod download

# 将代码复制到容器中
COPY ../.. .

# 将我们的代码编译成二进制可执行文件 GoDddDemo
RUN go build -o GoDddDemo .

###################
# 接下来创建一个小镜像
###################
FROM busybox
COPY --from=builder /build/GoDddDemo /GoDddDemo

# 拷贝配置文件到容器中
ENV WORKDIR  /app
COPY ../../manifest/config $WORKDIR/manifest/config
ADD ../../manifest/i18n     $WORKDIR/i18n

# 暴露端口
EXPOSE 21800
EXPOSE 21801
EXPOSE 21802

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

## 运行golang程序的命令
WORKDIR $WORKDIR
ENTRYPOINT ["/GoDddDemo", "-env", "debug"]