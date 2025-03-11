package main

import (
	"fmt"
	"net/http"
	"poetry-realtime/config"
	"poetry-realtime/server"
)

func main() {
	config.LoadConfig() // Load environment variables

	hub := server.NewHub()
	go hub.Run()

	port := ":" + config.Port
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWs(hub, w, r)
	})

	fmt.Println("WebSocket server started on", port)
	http.ListenAndServe(port, nil)
}
