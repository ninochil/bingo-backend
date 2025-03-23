# db-serverのDockerfile
FROM golang:1.24 AS builder

WORKDIR /app

# Go modulesのコピー
COPY ws-server/go.mod ws-server/go.sum ./
RUN go mod tidy

# ソースコードをコピー
COPY ws-server/. ./

# Airをインストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# バイナリをビルド
RUN go build -o bin/ws-server .

CMD ["air"]