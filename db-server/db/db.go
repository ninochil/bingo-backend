package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" 
)

// データベース接続を保持するグローバル変数
var DB *sql.DB

// Connect はデータベースへの接続を初期化
func Connect() error {
	var err error
	DB, err = sql.Open("mysql", "root:password@tcp(db:3306)/bingodon?parseTime=true") 
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("データベース接続成功")
	return nil
}