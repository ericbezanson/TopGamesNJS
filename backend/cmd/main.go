package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"topgames/api"
	"topgames/config"
)

func main() {
	config.LoadEnv() // Load environment variables

	http.HandleFunc("/games", api.GameHandler)

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
