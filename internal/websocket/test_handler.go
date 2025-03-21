package ws

import (
	"fmt"
	"net/http"
)

// NewWebSocketHandler returns a handler for WebSocket connections.
func NewWebSocketHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "WebSocket server is running")
		// WebSocket接続処理などをここで実装
	})
	return mux
}