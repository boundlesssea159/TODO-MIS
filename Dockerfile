# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.25-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用轻量级 alpine 镜像作为最终镜像
FROM alpine:latest

# 安装 ca-certificates 包，用于处理 SSL 证书
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建器镜像中复制编译后的二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8080

# 运行应用程序
CMD ["./main"]