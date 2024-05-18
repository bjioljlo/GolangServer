# 使用官方Golang鏡像作為基礎鏡像
FROM golang:latest

# 設置工作目錄
WORKDIR /GolangServer

# 將Go模組文件複製到容器中
COPY go.mod .
COPY go.sum .

# 下載所有依賴項
RUN go mod download

# 將源代碼複製到容器中
COPY . .

# 構建應用
RUN go build -o main .

EXPOSE 8080/tcp

# 啟動應用
CMD ["./main"]