package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"topgames/api"
	"topgames/config"
)

func main() {
	config.LoadEnv() // Load environment variables

	slog.Info("Hello World - App Start")
	http.HandleFunc("/games", api.GameHandler)
	http.HandleFunc("/gamedetail/", api.GameDetailHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)

	// FIX: Properly handle ListenAndServe
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
