package igdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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
	Cover            struct {
		ImageID string `json:"image_id"`
	} `json:"cover"` // 👈 Fix: Capture nested cover.image_id
	SimilarGames []int  `json:"similar_games"`
	Slug         string `json:"slug"`
	URL          string `json:"url"`
	CoverURL     string `json:"cover_url"` // Store full image URL
}

type PopularityData struct {
	GameID         int     `json:"game_id"`
	Value          float64 `json:"value"`
	PopularityType int     `json:"popularity_type"`
}

func FetchGames(accessToken string) ([]Game, error) {
	// Fetch popular game IDs
	popularGames, err := FetchMostPlayedGames(accessToken)
	if err != nil {
		return nil, err
	}

	// Format game IDs for IGDB query
	var gameIDs []string
	for _, game := range popularGames {
		gameIDs = append(gameIDs, fmt.Sprintf("%d", game.GameID))
	}
	gameIDString := fmt.Sprintf("(%s)", strings.Join(gameIDs, ", "))

	// Request detailed game info including cover.image_id
	query := fmt.Sprintf(`fields id, name, summary, genres, platforms, first_release_date, 
        aggregated_rating, rating, total_rating, screenshots, cover.image_id, similar_games, slug, url; 
        where id = %s;`, gameIDString)

	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/games", bytes.NewBufferString(query))
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

	// Construct full cover image URLs
	for i, game := range games {
		if game.Cover.ImageID != "" {
			games[i].CoverURL = GetCoverImageURL(game.Cover.ImageID, "cover_big")
		}
	}

	return games, nil
}

// FetchMostPlayedGames retrieves the top 10 most played games
func FetchMostPlayedGames(accessToken string) ([]PopularityData, error) {
	query := "fields game_id,value,popularity_type; sort value desc; limit 10; where popularity_type = 4;"

	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/popularity_primitives", bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	// Set headers for IGDB authentication
	req.Header.Set("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode response JSON into slice of PopularityData
	var games []PopularityData
	if err := json.NewDecoder(resp.Body).Decode(&games); err != nil {
		return nil, err
	}

	return games, nil
}

func GetCoverImageURL(imageID string, size string) string {
	return fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_%s/%s.jpg", size, imageID)
}
