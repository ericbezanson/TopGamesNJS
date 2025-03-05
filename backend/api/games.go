package api

import (
	"encoding/json"
	"net/http"
	"topgames/internal/igdb"
)

func GameHandler(w http.ResponseWriter, r *http.Request) {
	token, err := igdb.GetAccessToken()
	if err != nil {
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		return
	}

	games, err := igdb.FetchGames(token)
	if err != nil {
		http.Error(w, "Failed to fetch games", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}
