package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ninochil/bingo-backend/db"
)

// GetBingoCard ビンゴカード情報の取得 (ホストごとの)
func GetBingoCardHost(w http.ResponseWriter, r *http.Request) {
	hostID := r.URL.Query().Get("host_id")

	query := `SELECT bingo_card_id, host_id FROM Bingo_Card WHERE host_id = ?`
	rows, err := db.DB.Query(query, hostID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching bingo cards: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bingoCards []db.BingoCard
	for rows.Next() {
		var bingoCard db.BingoCard
		if err := rows.Scan(&bingoCard.BingoCardID, &bingoCard.HostID); err != nil {
			http.Error(w, fmt.Sprintf("Error reading bingo cards: %v", err), http.StatusInternalServerError)
			return
		}
		bingoCards = append(bingoCards, bingoCard)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"select_host_bingo_card": bingoCards})
}

// GetBingoCardPlayer プレイヤーごとのビンゴカード情報の取得
func GetBingoCardPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("player_id")


	query := `SELECT bingo_card_id, player_id FROM Bingo_Card WHERE player_id = ?`
	rows, err := db.DB.Query(query, playerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching bingo cards: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bingoCards []db.BingoCard
	for rows.Next() {
		var bingoCard db.BingoCard
		if err := rows.Scan(&bingoCard.BingoCardID, &bingoCard.PlayerID); err != nil {
			http.Error(w, fmt.Sprintf("Error reading bingo cards: %v", err), http.StatusInternalServerError)
			return
		}
		bingoCards = append(bingoCards, bingoCard)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"select_player_bingo_card": bingoCards})
}

// CreateBingoCard ビンゴカード情報の登録
func CreateBingoCard(w http.ResponseWriter, r *http.Request) {
	var bingoCard db.BingoCard
	if err := json.NewDecoder(r.Body).Decode(&bingoCard); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding bingo card data: %v", err), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Bingo_Card (bingo_card_id, host_id, player_id) VALUES (?, ?, ?)`
	_, err := db.DB.Exec(query, bingoCard.BingoCardID, bingoCard.HostID, bingoCard.PlayerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting bingo card data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}

// UpdateBingoCard ビンゴカード情報の更新
func UpdateBingoCard(w http.ResponseWriter, r *http.Request) {
	var bingoCard db.BingoCard
	if err := json.NewDecoder(r.Body).Decode(&bingoCard); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding bingo card data: %v", err), http.StatusBadRequest)
		return
	}

	var query string
	var params []interface{}

	// `bingo_card_id` が提供されている場合はそれを基に更新
	if bingoCard.BingoCardID != "" {
		query = `UPDATE Bingo_Card SET host_id = ?, player_id = ? WHERE bingo_card_id = ?`
		params = append(params, bingoCard.HostID, bingoCard.PlayerID, bingoCard.BingoCardID)
	} else if bingoCard.PlayerID != "" {
		// `player_id` が提供されている場合はそれを基に更新
		query = `UPDATE Bingo_Card SET host_id = ? WHERE player_id = ?`
		params = append(params, bingoCard.HostID, bingoCard.PlayerID)
	} else {
		http.Error(w, "Either bingo_card_id or player_id must be provided for update", http.StatusBadRequest)
		return
	}

	// クエリ実行
	_, err := db.DB.Exec(query, params...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating bingo card data: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功応答
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}