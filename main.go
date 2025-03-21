package main

import (
	"net/http"
	"sync"

	"github.com/ninochil/bingo-backend/internal/db"
	"github.com/ninochil/bingo-backend/internal/logger"
	"github.com/ninochil/bingo-backend/internal/websocket"
)

func main() {
	// ロガーの初期化
	logger.InitLogger()

	// WaitGroupで並列起動（DBサーバーとWebSocketサーバーの起動）
	var wg sync.WaitGroup
	wg.Add(2)

	// 5001ポートでDBサーバー起動
	go func() {
		defer wg.Done()
		logger.Info("Starting DB server on port 5001")
		err := db.StartDBServer(":5001")
		if err != nil {
			logger.Error(err)
		}
	}()

	// 5002ポートでWebSocketサーバー起動
	go func() {
		defer wg.Done()
		logger.Info("Starting WebSocket server on port 5002")
		err := http.ListenAndServe(":5002", ws.NewWebSocketHandler())
		if err != nil {
			logger.Error(err)
		}
	}()

	// 両方のサーバーが終了するまで待つ
	wg.Wait()
	logger.Info("All servers stopped")
}