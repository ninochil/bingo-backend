package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"sync"
)

// WebSocket接続を管理するための構造体
type WebSocketServer struct {
	upgrader websocket.Upgrader
}

// NewWebSocketServer はWebSocketサーバーのインスタンスを生成します。
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }, // セキュリティ設定 (オリジンチェックを無効化)
		},
	}
}

// HandleWebSocketはWebSocket接続を受け付け、メッセージを処理します。
func (ws *WebSocketServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	for {
		// クライアントからのメッセージを受信
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// メッセージをそのまま返す (エコーサーバー)
		if err := conn.WriteMessage(msgType, p); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

// StartServerはWebSocketサーバーを指定されたポートで起動します。
func (ws *WebSocketServer) StartServer(address string) {
	http.HandleFunc("/ws", ws.HandleWebSocket)

	// サーバー起動
	fmt.Println("WebSocket server started on", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func main() {
	// WebSocketサーバーのインスタンス作成
	wsServer := NewWebSocketServer()

	// WebSocketサーバーを5002番ポートで起動
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		wsServer.StartServer(":5002")
	}()

	// 並列処理でサーバーを待機
	wg.Wait()
	fmt.Println("WebSocket server stopped")
}