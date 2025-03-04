package igdb

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type Game struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Summary          string  `json:"summary"`
	Genres           []int   `json:"genres"`
	Platforms        []int   `json:"platforms"`
	FirstReleaseDate int64   `json:"first_release_date"`
	AggregatedRating float64 `json:"aggregated_rating"`
	Rating           float64 `json:"rating"`
	TotalRating      float64 `json:"total_rating"`
	Screenshots      []int   `json:"screenshots"`
	Cover            int     `json:"cover"`
	SimilarGames     []int   `json:"similar_games"`
	Slug             string  `json:"slug"`
	URL              string  `json:"url"`
}

func FetchGames(accessToken string) ([]Game, error) {
	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/games", bytes.NewBufferString("fields *; where id = 279661;")) // !TODO: single game is hardcoded, will add support to loop through list of top games
	if err != nil {
		return nil, err
	}

	req.Header.Set("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var games []Game
	if err := json.NewDecoder(resp.Body).Decode(&games); err != nil {
		return nil, err
	}

	return games, nil
}
