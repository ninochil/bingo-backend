package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/ninochil/bingo-backend/db/handler"
	"net/http"
)

// Server はAPIサーバーの構造体
type Server struct {
	router *gin.Engine
	db     *sql.DB
}

// NewServer はAPIサーバーの新しいインスタンスを作成
func NewServer(db *sql.DB) *Server {
	router := gin.Default() // Ginを使ったルータ作成
	server := &Server{
		router: router,
		db:     db,
	}

	// CORS設定
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // 全てのオリジンを許可
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // OPTIONSリクエストに対して204を返す
			return
		}
		c.Next()
	})

	// ルート設定
	server.routes()

	return server
}

//ルーティングの設定
func (s *Server) routes() {
	// ホスト関連のAPI
	s.router.GET("/api/host/select", s.wrapHandler(handler.GetHost))               // ホスト情報の取得
	s.router.POST("/api/host/insert", s.wrapHandler(handler.CreateHost))           // ホスト情報の登録
	s.router.DELETE("/api/host/delete", s.wrapHandler(handler.DeleteHost))         // ホスト情報の削除

	// プレイヤー関連のAPI
	s.router.GET("/api/player/select", s.wrapHandler(handler.GetPlayer))           // プレイヤー情報の取得
	s.router.POST("/api/player/insert", s.wrapHandler(handler.CreatePlayer))       // プレイヤー情報の登録

	// 質問関連のAPI
	s.router.GET("/api/question/select", s.wrapHandler(handler.GetQuestion))       // 質問内容の取得
	s.router.POST("/api/question/insert", s.wrapHandler(handler.CreateQuestion))   // 質問内容の登録
	s.router.PUT("/api/question/update", s.wrapHandler(handler.UpdateQuestion))    // 質問内容の更新

	// ビンゴカード関連のAPI
	s.router.GET("/api/host/bingo_card/select", s.wrapHandler(handler.GetBingoCardHost))   // ホスト毎のビンゴカード情報の取得
	s.router.GET("/api/player/bingo_card/select", s.wrapHandler(handler.GetBingoCardPlayer)) // プレイヤー毎のビンゴカード情報の取得
	s.router.POST("/api/bingo_card/insert", s.wrapHandler(handler.CreateBingoCard))         // ビンゴカード情報の登録
	s.router.PUT("/api/bingo_card/update", s.wrapHandler(handler.UpdateBingoCard))          // ビンゴカード情報の更新

	// ビンゴカードセル状態関連のAPI
	s.router.GET("/api/bingo_card_cells_status/select", s.wrapHandler(handler.GetBingoCardCellsStatus))  // ビンゴカードセル状態の取得
	s.router.POST("/api/bingo_card_cells_status/insert", s.wrapHandler(handler.CreateBingoCardCellsStatus)) // ビンゴカードセル状態の作成
	s.router.PUT("/api/bingo_card_cells_status/update", s.wrapHandler(handler.UpdateBingoCardCellsStatus)) // ビンゴカードセル状態の更新

	// ホスト毎のQuestion_Usage関連のAPI
	s.router.GET("/api/host/question_usage/select", s.wrapHandler(handler.GetHostQuestionUsage))  // ホスト毎のQuestion_Usage情報の取得
	s.router.POST("/api/question_usage/insert", s.wrapHandler(handler.CreateQuestionUsage))     // Question_Usage情報の登録

	// プレイヤー毎のQuestion_Usage関連のAPI
	s.router.GET("/api/player/question_usage/select", s.wrapHandler(handler.GetPlayerQuestionUsage)) // プレイヤー毎のQuestion_Usage情報の取得
	s.router.PUT("/api/question_usage/update", s.wrapHandler(handler.UpdateQuestionUsage))           // Question_Usage情報の更新
}

// wrapHandler は通常の http.HandlerFunc を gin.HandlerFunc にラップ
func (s *Server) wrapHandler(fn func(w http.ResponseWriter, r *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c.Writer, c.Request) // Ginのコンテキストを http.HandlerFunc に渡す
	}
}

// サーバーを指定されたアドレスで起動
func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}