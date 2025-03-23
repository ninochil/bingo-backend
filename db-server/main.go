package main

import (
	"github.com/ninochil/bingo-backend/api"
	"github.com/ninochil/bingo-backend/db"
	"log"
)

func main() {
	// データベース接続
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// マイグレーション
	db.Migrate(db.DB)


	// APIサーバー作成
	server := api.NewServer(db.DB)

	// サーバー起動
	err = server.Start(":5001")
	if err != nil {
		log.Fatal(err)
	}
}