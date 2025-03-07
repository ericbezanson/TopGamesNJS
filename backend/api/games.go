package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
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

func GameDetailHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("GameDetailHandler invoked")
	// Extract the path from the URL
	path := r.URL.Path

	// Split the path into segments
	segments := strings.Split(strings.Trim(path, "/"), "/")

	// Check if the path has the expected format: /gamedetail/{id}
	if len(segments) != 2 || segments[0] != "gamedetail" {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	// Extract and convert the game ID
	gameID, err := strconv.Atoi(segments[1])

	slog.Info("GameID", gameID)

	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	token, err := igdb.GetAccessToken()
	if err != nil {
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		return
	}

	// Fetch the game details by ID
	game, err := igdb.FetchGameByID(token, gameID)
	if err != nil {
		http.Error(w, "Failed to fetch game details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
