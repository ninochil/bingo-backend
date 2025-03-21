package db

import (
	"fmt"
	"net/http"
)

// StartDBServer starts a simple HTTP server to simulate DB server behavior.
func StartDBServer(addr string) error {
	http.HandleFunc("/db", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "DB server is running")
	})
	return http.ListenAndServe(addr, nil)
}