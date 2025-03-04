package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"topgames/internal/igdb"
)

func GameHandler(w http.ResponseWriter, r *http.Request) {
	token, err := igdb.GetAccessToken()
	if err != nil {
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		return
	}
	fmt.Println("TOKEN", token)
	games, err := igdb.FetchGames(token)
	if err != nil {
		http.Error(w, "Failed to fetch games", http.StatusInternalServerError)
		return
	}
	fmt.Println("GAMES", games)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}
