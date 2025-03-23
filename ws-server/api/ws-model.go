package api

import (
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type User struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

type Room struct {
	RoomID    string
	Users     map[string]*User
	GameState *GameState
}

type GameState struct {
	RouletteNumber int              // ルーレットの番号
	Votes          map[string]int   // ユーザーごとの投票番号
	BingoNumbers   []int            // ビンゴの当選番号
}

type Client struct {
	Conn      *websocket.Conn
	UserID    string
	UserName  string
	Numbers   []int
	Questions []string
}

type Message struct {
	Type      string   `json:"type"`
	UserID    string   `json:"userId,omitempty"`
	UserName  string   `json:"userName,omitempty"`
	Number    int      `json:"number,omitempty"`
	Users     []string `json:"users,omitempty"`
	Numbers   []int    `json:"numbers,omitempty"`
	Questions []string `json:"questions,omitempty"`
	Question  string   `json:"question,omitempty"`
	Old       string   `json:"old,omitempty"`
	New       string   `json:"new,omitempty"`
	Result    string   `json:"result,omitempty"`
	VotedUser string   `json:"votedUser,omitempty"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var mu sync.Mutex

var users []User
var userClients []*websocket.Conn

var Clients = make(map[*Client]bool)
var Broadcast = make(chan Message)
var votes = make(map[string]string)
var currentQuestion string
var matchedUsers []string

var QuestionBank = []string{
	"質問1", "質問2", "質問3", "質問4", "質問5", "質問6", "質問7", "質問8", "質問9",
	"質問10", "質問11", "質問12", "質問13", "質問14", "質問15", "質問16", "質問17", "質問18",
	"質問19", "質問20", "質問21", "質問22", "質問23", "質問24", "質問25", "質問26", "質問27",
}

func GenerateQuestionLabels() []string {
	indices := rand.Perm(len(QuestionBank))[:9]
	result := make([]string, 9)
	for i, idx := range indices {
		result[i] = QuestionBank[idx]
	}
	return result
}