package yt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// Video represents the structure of a video fetched from YouTube API

type Video struct {
	Title        string
	Description  string
	PublishedAt  time.Time
	ThumbnailURL string // Change this to string
}

type APIResponse struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			PublishedAt string `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

func FetchVideos(apiKey, query string) ([]Video, error) {
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=50&q=%s&type=video&order=date&key=%s", query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	var videos []Video
	for _, item := range apiResp.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			log.Println("Error parsing date:", err)
			continue
		}

		video := Video{
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishedAt:  publishedAt,
			ThumbnailURL: item.Snippet.Thumbnails.Default.URL, // This is now a string
		}

		videos = append(videos, video)
	}

	return videos, nil
}

// StoreVideos stores the fetched videos in the database
// StoreVideos stores the fetched videos in the database
func StoreVideos(db *gorm.DB, videos []Video) {
	for _, video := range videos {
		if err := db.Create(&video).Error; err != nil {
			log.Printf("Error storing video: %v", err)
		}
		log.Printf("Stored video: %v", video)
	}
}

// FetchAndStoreVideos fetches the videos using multiple API keys and stores them in the database
func FetchAndStoreVideos(db *gorm.DB, apiKeys []string, query string, interval time.Duration) {
	var currentAPIKey int

	for {
		videos, err := FetchVideos(apiKeys[currentAPIKey], query)
		if err != nil {
			log.Printf("Error fetching videos with API key %d: %v", currentAPIKey, err)
			currentAPIKey = (currentAPIKey + 1) % len(apiKeys) // Rotate API key
			continue
		}

		StoreVideos(db, videos)

		time.Sleep(interval) // Sleep for the configured interval
	}
}
