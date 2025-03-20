DOCKER_COMPOSE = docker compose
TARGET ?= db-server

# DBサーバーのビルド＆起動
build-db:
	$(DOCKER_COMPOSE) build --build-arg TARGET=db-server

run-db: build-db
	$(DOCKER_COMPOSE) up -d db-server

# WebSocketサーバーのビルド＆起動
build-ws:
	$(DOCKER_COMPOSE) build --build-arg TARGET=ws-server

run-ws: build-ws
	$(DOCKER_COMPOSE) up -d ws-server

# サーバー停止
stop:
	$(DOCKER_COMPOSE) down

# 不要なコンテナやイメージ削除
clean:
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans