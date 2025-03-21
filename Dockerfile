FROM golang:1.24 as builder

# 作業ディレクトリを作成
WORKDIR /app

# go.modとgo.sumをコピー
COPY go.mod go.sum ./

# 依存関係を解決
RUN go mod tidy

# ソースコードをコンテナにコピー
COPY . .

# 安定バージョンの air を明示的にインストール
RUN go install github.com/cosmtrek/air@v1.43.0

# air を使ってホットリロード実行
CMD ["air"]