package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ninochil/bingo-backend/db"
)

// GetPlayer プレイヤー情報の取得
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("player_id")


	query := `SELECT player_id, name FROM Player WHERE player_id = ?`
	row := db.DB.QueryRow(query, playerID)

	var player db.Player
	err := row.Scan(&player.PlayerID, &player.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching player: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"select_player": player})
}

// CreatePlayer プレイヤー情報の登録
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player db.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding player data: %v", err), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Player (player_id, name) VALUES (?, ?)`
	_, err := db.DB.Exec(query, player.PlayerID, player.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting player data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}