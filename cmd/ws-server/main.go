package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ninochil/bingo-backend/internal"
)

func main() {
    r := mux.NewRouter()
    // r.HandleFunc("/ws/room", internal.HandleRoom)
    r.HandleFunc("/ws/user", internal.HandleUser)
    r.HandleFunc("/ws/game", internal.HandleGame)

    go internal.HandleMessages()

    port := ":5002"
    fmt.Println("サーバーを",port,"で起動しました。")
    log.Fatal(http.ListenAndServe(port, r))
}