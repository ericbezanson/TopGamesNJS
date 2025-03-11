package igdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Screenshot struct {
	ImageID string `json:"image_id"`
	URL     string `json:"url,omitempty"`
}

type TopGame struct {
	ID       int    `json:"id"`
	CoverURL string `json:"cover_url"` // Store full image URL
	Cover    struct {
		ID      int    `json:"id"`
		ImageID string `json:"image_id"`
	} `json:"cover"`
	Name string `json:"name"`
}

type GameDetail struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Summary          string       `json:"summary"`
	Genres           []int        `json:"genres"`
	Platforms        []int        `json:"platforms"`
	FirstReleaseDate int64        `json:"first_release_date"`
	AggregatedRating float64      `json:"aggregated_rating"`
	Rating           float64      `json:"rating"`
	TotalRating      float64      `json:"total_rating"`
	Screenshots      []Screenshot `json:"screenshots,omitempty"`
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

func FetchGameByID(accessToken string, gameID int) (*GameDetail, error) {
	// Construct the query to fetch a single game's details
	query := fmt.Sprintf(`fields id, name, summary, genres, platforms, first_release_date, 
    aggregated_rating, rating, total_rating, screenshots.image_id, cover.image_id, similar_games, slug, url; 
    where id = %d;`, gameID)

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

	// Log the raw response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Raw Response:", string(body))

	// Parse the JSON response
	var games []GameDetail
	if err := json.Unmarshal(body, &games); err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, fmt.Errorf("game not found")
	}

	game := games[0]
	// Construct the full cover image URL
	if games[0].Cover.ImageID != "" {
		games[0].CoverURL = GetImageURL(games[0].Cover.ImageID, "cover_big")
	}
	// Construct the full screenshot URLs
	for i, screenshot := range game.Screenshots {
		if screenshot.ImageID != "" {
			game.Screenshots[i].URL = GetImageURL(screenshot.ImageID, "1080p")
		}
	}
	return &games[0], nil
}

func FetchTopGames(accessToken string) ([]TopGame, error) {
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
	query := fmt.Sprintf(`fields id, name, cover.image_id; where id = %s;`, gameIDString)

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

	// Read and parse the JSON response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var games []TopGame
	if err := json.Unmarshal(body, &games); err != nil {
		return nil, err
	}

	// Construct full cover image URLs
	for i, game := range games {
		if game.Cover.ImageID != "" {
			games[i].CoverURL = GetImageURL(game.Cover.ImageID, "cover_big")
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

func GetImageURL(imageID, size string) string {
	return fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_%s/%s.jpg", size, imageID)
}
