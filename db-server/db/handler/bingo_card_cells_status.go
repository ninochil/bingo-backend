package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ninochil/bingo-backend/db"
)

// GetBingoCardCellsStatus ビンゴカードセル状態の取得
func GetBingoCardCellsStatus(w http.ResponseWriter, r *http.Request) {
	bingoCardID := r.URL.Query().Get("bingo_card_id")

	query := `SELECT bingo_card_id, question_id, position_x, position_y, is_checked FROM Bingo_Card_Cells_Status WHERE bingo_card_id = ?`
	rows, err := db.DB.Query(query, bingoCardID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching bingo card cells status: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var statusList []db.BingoCardCellsStatus
	for rows.Next() {
		var status db.BingoCardCellsStatus
		if err := rows.Scan(&status.BingoCardID, &status.QuestionID, &status.PositionX, &status.PositionY, &status.IsChecked); err != nil {
			http.Error(w, fmt.Sprintf("Error reading bingo card cells status: %v", err), http.StatusInternalServerError)
			return
		}
		statusList = append(statusList, status)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"select_bingo_card_cells_status": statusList})
}

// CreateBingoCardCellsStatus ビンゴカードセル状態の作成
func CreateBingoCardCellsStatus(w http.ResponseWriter, r *http.Request) {
	var status db.BingoCardCellsStatus
	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding bingo card cells status: %v", err), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Bingo_Card_Cells_Status (bingo_card_id, question_id, position_x, position_y) VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, status.BingoCardID, status.QuestionID, status.PositionX, status.PositionY)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting bingo card cells status: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}

// UpdateBingoCardCellsStatus ビンゴカードセル状態の更新
func UpdateBingoCardCellsStatus(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータから値を取得
	bingoCardID := r.URL.Query().Get("bingo_card_id")
	questionID := r.URL.Query().Get("question_id")
	positionX := r.URL.Query().Get("position_x")
	positionY := r.URL.Query().Get("position_y")

	// リクエストボディから JSON をデコード
	var status db.BingoCardCellsStatus
	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding bingo card cells status: %v", err), http.StatusBadRequest)
		return
	}

	// 更新クエリの開始部分を準備
	query := `UPDATE Bingo_Card_Cells_Status SET `
	params := []interface{}{}

	// isChecked がリクエストボディに含まれていれば、それを更新する
	if status.IsChecked != nil {
		query += `is_checked = ?`
		params = append(params, *status.IsChecked)  // ポインタを解 dereference して使用
	}

	// isChecked が含まれていなければ、SET句をそのまま終了
	if len(params) == 0 {
		http.Error(w, "No valid fields to update", http.StatusBadRequest)
		return
	}

	// WHERE 句が必ず必要なので追加
	query += ` WHERE bingo_card_id = ? AND question_id = ? AND position_x = ? AND position_y = ?`
	params = append(params, bingoCardID, questionID, positionX, positionY)

	// クエリ実行
	_, err := db.DB.Exec(query, params...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating bingo card cells status: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}