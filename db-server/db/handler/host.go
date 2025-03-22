package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ninochil/bingo-backend/db"
)

// GetHost ホスト情報の取得
func GetHost(w http.ResponseWriter, r *http.Request) {
	hostID := r.URL.Query().Get("host_id")

	query := `SELECT host_id, room_code FROM Host WHERE host_id = ?`
	row := db.DB.QueryRow(query, hostID)

	var host db.Host
	err := row.Scan(&host.HostID, &host.RoomCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching host: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"select_host": host})
}

// CreateHost ホスト情報の登録
func CreateHost(w http.ResponseWriter, r *http.Request) {
	var host db.Host
	if err := json.NewDecoder(r.Body).Decode(&host); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding host data: %v", err), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Host (host_id, room_code) VALUES (?, ?)`
	_, err := db.DB.Exec(query, host.HostID, host.RoomCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting host data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}

// DeleteHost ホスト情報の削除
func DeleteHost(w http.ResponseWriter, r *http.Request) {
	hostID := r.URL.Query().Get("host_id")

	query := `DELETE FROM Host WHERE host_id = ?`
	_, err := db.DB.Exec(query, hostID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting host: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}