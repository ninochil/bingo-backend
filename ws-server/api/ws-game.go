package api

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
			votes = make(map[string]string)
			rouletteIndex := rand.Intn(len(QuestionBank))
			currentQuestion = QuestionBank[rouletteIndex]
			matchedUsers = []string{}

			mu.Lock()
			for c := range Clients {
				for _, q := range c.Questions {
					if q == currentQuestion {
						matchedUsers = append(matchedUsers, c.UserName)
						break
					}
				}
			}
			mu.Unlock()

			Broadcast <- Message{
				Type:     "rouletteInfo",
				Question: currentQuestion,
				Users:    matchedUsers,
			}

		} else if msg.Type == "vote" {
			votes[msg.UserID] = msg.VotedUser
			if len(votes) == len(Clients)-len(matchedUsers)-1 {
				Broadcast <- Message{Type: "votesCompleted"}
			}

		} else if msg.Type == "showResult" {
			voteCounts := make(map[string]int)
			for _, voted := range votes {
				voteCounts[voted]++
			}

			topUser := ""
			maxVotes := 0
			for name, count := range voteCounts {
				if count > maxVotes {
					topUser = name
					maxVotes = count
				}
			}

			mu.Lock()
			for c := range Clients {
				if c.UserName == topUser {
					c.Conn.WriteJSON(Message{
						Type:     "markQuestion",
						Question: currentQuestion,
					})
				} else {
					for i, q := range c.Questions {
						if q == currentQuestion {
							newQ := GetRandomNewQuestion(c.Questions)
							c.Questions[i] = newQ
							c.Conn.WriteJSON(Message{
								Type: "updateQuestion",
								Old:  currentQuestion,
								New:  newQ,
							})
							break
						}
					}
				}
			}
			mu.Unlock()

			Broadcast <- Message{
				Type:   "voteResult",
				Result: topUser,
			}

			votes = make(map[string]string)
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

func GetRandomNewQuestion(current []string) string {
	existing := make(map[string]bool)
	for _, q := range current {
		existing[q] = true
	}

	candidates := []string{}
	for _, q := range QuestionBank {
		if !existing[q] {
			candidates = append(candidates, q)
		}
	}

	if len(candidates) == 0 {
		return "(質問なし)"
	}
	return candidates[rand.Intn(len(candidates))]
}