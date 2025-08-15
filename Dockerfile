# Dockerfile

# ---- Stage 1: Build ----
# 使用官方的 Go 镜像作为构建环境
# 我们选择一个具体的版本以保证构建的可重现性
FROM golang:1.24.3-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
# 这样做可以利用 Docker 的层缓存机制。只要这两个文件没变，就不需要重新下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制所有源代码
COPY . .

# 编译应用
# CGO_ENABLED=0: 禁用 CGO，允许我们静态链接，生成一个纯 Go 的可执行文件
# -ldflags="-w -s": 减小可执行文件的大小
# -o /app/server: 指定输出文件名为 server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server ./cmd/server/main.go

# ---- Stage 2: Production ----
# 使用一个极小的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /app/server /app/server

# 复制配置文件目录
# 注意：这是一种方式，但更好的方式是通过挂载卷或配置中心
# 这里为了简化，我们将其打包进镜像
COPY configs /app/configs

# 暴露应用监听的端口
EXPOSE 8080

# 容器启动时执行的命令
CMD ["/app/server"]