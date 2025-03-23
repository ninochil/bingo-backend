package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	userClients = append(userClients, conn)
	mu.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)

			mu.Lock()
			for i, client := range userClients {
				if client == conn {
					userClients = append(userClients[:i], userClients[i+1:]...)
					break
				}
			}
			mu.Unlock()
			break
		}

		var data map[string]interface{}
		json.Unmarshal(msg, &data)

		if data["type"] == "join" {
			mu.Lock()
			user := User{
				UserID:   uuid.New().String(),
				UserName: data["userName"].(string),
			}
			users = append(users, user)
			mu.Unlock()

			response := map[string]interface{}{
				"type":   "userId",
				"userId": user.UserID,
			}
			res, _ := json.Marshal(response)
			conn.WriteMessage(websocket.TextMessage, res)

			for _, client := range userClients {
				client.WriteJSON(map[string]interface{}{
					"type": "newUser",
					"user": user,
				})
			}
		}

		if data["type"] == "register" {
			mu.Lock()
			response := map[string]interface{}{
				"type":  "userList",
				"users": users,
			}
			res, _ := json.Marshal(response)
			mu.Unlock()

			conn.WriteMessage(websocket.TextMessage, res)
		}

		if data["type"] == "cancel" {
			userId, ok := data["userId"].(string)
			if !ok {
				fmt.Println("Error: userId is not a string")
				continue
			}
			mu.Lock()
			for i, user := range users {
				if user.UserID == userId {
					users = append(users[:i], users[i+1:]...)
					break
				}
			}
			mu.Unlock()

			for _, client := range userClients {
				client.WriteJSON(map[string]interface{}{
					"type":   "userLeft",
					"userId": userId,
				})
			}
		}

		if data["type"] == "gameStart" {
			for _, client := range userClients {
				client.WriteJSON(map[string]interface{}{
					"type": "gameStart",
				})
			}
		}
	}
}