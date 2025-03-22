package internal

import (
	"log"
	"math/rand"
	"net/http"
)

func HandleGame(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	client := &Client{Conn: conn}
	mu.Lock()
	Clients[client] = true
	mu.Unlock()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		if msg.Type == "joinGame" {
			client.UserID = msg.UserID
			client.UserName = msg.UserName
			client.Questions = GenerateQuestionLabels()

			client.Conn.WriteJSON(Message{
				Type:      "bingoQuestions",
				Questions: client.Questions,
			})
		} else if msg.Type == "spinRoulette" {
			// ランダムな質問を1つ選ぶ
			rouletteIndex := rand.Intn(len(QuestionBank))
			targetQuestion := QuestionBank[rouletteIndex]

			matchedUsers := []string{}

			mu.Lock()
			for c := range Clients {
				for _, q := range c.Questions {
					if q == targetQuestion {
						matchedUsers = append(matchedUsers, c.UserName)
						c.Conn.WriteJSON(Message{
							Type:     "markQuestion",
							Question: q,
						})
						break
					}
				}
			}
			mu.Unlock()

			Broadcast <- Message{
				Type:     "rouletteInfo",
				Question: targetQuestion,
				Users:    matchedUsers,
			}
		}
	}

	mu.Lock()
	delete(Clients, client)
	mu.Unlock()
}

func HandleMessages() {
	for {
		msg := <-Broadcast
		mu.Lock()
		for client := range Clients {
			client.Conn.WriteJSON(msg)
		}
		mu.Unlock()
	}
}