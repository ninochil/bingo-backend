package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Migrate はデータベースに必要なテーブルを作成する関数です。
func Migrate(db *sql.DB) {
	// Hostテーブルの作成
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Host (
			host_id VARCHAR(255) PRIMARY KEY,
			room_code VARCHAR(255) NOT NULL UNIQUE
		);
	`)
	if err != nil {
		log.Fatalf("Error creating Host table: %v", err)
	}
	fmt.Println("Host table created successfully")

	// Playerテーブルの作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Player (
			player_id VARCHAR(255) PRIMARY KEY,
			host_id VARCHAR(255),
			name VARCHAR(255) NOT NULL,
			FOREIGN KEY (host_id) REFERENCES Host(host_id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		log.Fatalf("Error creating Player table: %v", err)
	}
	fmt.Println("Player table created successfully")

	// Questionテーブルの作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Question (
			question_id VARCHAR(255) PRIMARY KEY,
			question VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("Error creating Question table: %v", err)
	}
	fmt.Println("Question table created successfully")

	// Bingo_Cardテーブルの作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Bingo_Card (
			bingo_card_id VARCHAR(255) PRIMARY KEY,
			host_id VARCHAR(255),
			player_id VARCHAR(255),
			is_bingo BOOLEAN DEFAULT FALSE,
			FOREIGN KEY (host_id) REFERENCES Host(host_id) ON DELETE CASCADE,
			FOREIGN KEY (player_id) REFERENCES Player(player_id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		log.Fatalf("Error creating Bingo_Card table: %v", err)
	}
	fmt.Println("Bingo_Card table created successfully")

	// Bingo_Card_Cells_Statusテーブルの作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Bingo_Card_Cells_Status (
			bingo_card_id VARCHAR(255),
			question_id VARCHAR(255),
			position_x INT NOT NULL,
			position_y INT NOT NULL,
			is_checked BOOLEAN DEFAULT FALSE,
			PRIMARY KEY (bingo_card_id, question_id, position_x, position_y),
			FOREIGN KEY (bingo_card_id) REFERENCES Bingo_Card(bingo_card_id) ON DELETE CASCADE,
			FOREIGN KEY (question_id) REFERENCES Question(question_id)
		);
	`)
	if err != nil {
		log.Fatalf("Error creating Bingo_Card_Cells_Status table: %v", err)
	}
	fmt.Println("Bingo_Card_Cells_Status table created successfully")

	// Question_Usageテーブルの作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Question_Usage (
			question_id VARCHAR(255),
			host_id VARCHAR(255),
			player_id VARCHAR(255),
			number_of_uses INT DEFAULT 0,
			PRIMARY KEY (question_id, host_id, player_id),
			FOREIGN KEY (question_id) REFERENCES Question(question_id),
			FOREIGN KEY (host_id) REFERENCES Host(host_id) ON DELETE CASCADE,
			FOREIGN KEY (player_id) REFERENCES Player(player_id)
		);
	`)
	if err != nil {
		log.Fatalf("Error creating Question_Usage table: %v", err)
	}
	fmt.Println("Question_Usage table created successfully")
}