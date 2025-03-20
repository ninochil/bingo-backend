# Go 1.21 の公式イメージを使用
FROM golang:1.24 as builder

# 作業ディレクトリを作成
WORKDIR /app

# go.modとgo.sumを先にコピーしてキャッシュを効率化
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# ビルドターゲットを受け取ってビルド
ARG TARGET
RUN go build -o /app/app ./cmd/${TARGET}

# 実行用のイメージ（軽量化）
FROM gcr.io/distroless/base-debian12

# 必要なファイルをコピー
COPY --from=builder /app/app /app

# 実行するコマンド
CMD ["/app"]