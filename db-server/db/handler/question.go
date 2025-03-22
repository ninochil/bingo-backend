package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ninochil/bingo-backend/db"
)

// GetQuestion 質問情報を取得
func GetQuestion(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("question_id")

	query := `SELECT question_id, question FROM Question WHERE question_id = ?`
	row := db.DB.QueryRow(query, questionID)

	var question db.Question
	err := row.Scan(&question.QuestionID, &question.Question)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching question: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

// CreateQuestion 質問情報を作成
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question db.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding question data: %v", err), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Question (question_id, question) VALUES (?, ?)`
	_, err := db.DB.Exec(query, question.QuestionID, question.Question)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting question data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Question created successfully")
}

// UpdateQuestion 質問情報を更新
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("question_id")

	var question db.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding question data: %v", err), http.StatusBadRequest)
		return
	}

	query := `UPDATE Question SET question = ? WHERE question_id = ?`
	_, err := db.DB.Exec(query, question.Question, questionID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating question data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Question updated successfully")
}