package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ninochil/bingo-backend/db"
)

// GetHostQuestionUsage ホスト毎のQuestion_Usageテーブル情報の取得
func GetHostQuestionUsage(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("question_id")
	hostID := r.URL.Query().Get("host_id")

	query := `SELECT question_id, host_id, player_id, number_of_uses FROM Question_Usage WHERE question_id = ? AND host_id = ?`
	rows, err := db.DB.Query(query, questionID, hostID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching host question usage: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var usageList []db.QuestionUsage
	for rows.Next() {
		var usage db.QuestionUsage
		if err := rows.Scan(&usage.QuestionID, &usage.HostID, &usage.PlayerID, &usage.NumberOfUses); err != nil {
			http.Error(w, fmt.Sprintf("Error reading host question usage: %v", err), http.StatusInternalServerError)
			return
		}
		usageList = append(usageList, usage)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"select_host_question_usage": usageList})
}

// GetPlayerQuestionUsage プレイヤー毎のQuestion_Usageテーブル情報の取得
func GetPlayerQuestionUsage(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("question_id")
	playerID := r.URL.Query().Get("player_id")

	query := `SELECT question_id, host_id, player_id, number_of_uses FROM Question_Usage WHERE question_id = ? AND player_id = ?`
	rows, err := db.DB.Query(query, questionID, playerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching player question usage: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var usageList []db.QuestionUsage
	for rows.Next() {
		var usage db.QuestionUsage
		if err := rows.Scan(&usage.QuestionID, &usage.HostID, &usage.PlayerID, &usage.NumberOfUses); err != nil {
			http.Error(w, fmt.Sprintf("Error reading player question usage: %v", err), http.StatusInternalServerError)
			return
		}
		usageList = append(usageList, usage)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"select_player_question_usage": usageList})
}

// CreateQuestionUsage Question_Usageテーブル情報の登録
func CreateQuestionUsage(w http.ResponseWriter, r *http.Request) {
	var usage db.QuestionUsage
	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding question usage data: %v", err), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Question_Usage (question_id, host_id, player_id) VALUES (?, ?, ?)`
	_, err := db.DB.Exec(query, usage.QuestionID, usage.HostID, usage.PlayerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting question usage data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}

// UpdateQuestionUsage Question_Usageテーブル情報の更新
func UpdateQuestionUsage(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("question_id")
	playerID := r.URL.Query().Get("player_id")

	var usage db.QuestionUsage
	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding question usage data: %v", err), http.StatusBadRequest)
		return
	}

	query := `UPDATE Question_Usage SET number_of_uses = ? WHERE question_id = ? OR player_id = ?`
	_, err := db.DB.Exec(query, usage.NumberOfUses, questionID, playerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating question usage data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}