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
	// Construct the YouTube API request URL
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=50&q=%s&type=video&order=date&key=%s", query, apiKey)

	log.Printf("Making request to URL: %s", url)

	// Send the HTTP GET request to the YouTube API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch videos: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is not 200 (OK)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	// Decode the JSON response into the APIResponse struct
	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	// Log the response for debugging
	log.Printf("API Response: %+v", apiResp)

	// Iterate over the items in the API response and convert them into Video structs
	var videos []Video
	for _, item := range apiResp.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			log.Printf("Error parsing date for video %s: %v", item.Snippet.Title, err)
			continue
		}

		video := Video{
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishedAt:  publishedAt,
			ThumbnailURL: item.Snippet.Thumbnails.Default.URL,
		}

		log.Printf("Adding video: %s", video.Title)
		videos = append(videos, video)
	}

	// Log the number of videos fetched
	log.Printf("Fetched %d videos", len(videos))

	// Return the list of videos
	return videos, nil
}

// StoreVideos stores the fetched videos in the database
func StoreVideos(db *gorm.DB, videos []Video) {
	// Use a bulk insert operation to improve performance
	if err := db.Create(&videos).Error; err != nil {
		log.Printf("Failed to store videos: %v", err)
		return
	}

	log.Printf("Successfully stored %d videos.", len(videos))
}

// FetchAndStoreVideos fetches the videos using multiple API keys and stores them in the database
func FetchAndStoreVideos(db *gorm.DB, apiKeys []string, query string, interval time.Duration) {
	for {
		for _, apiKey := range apiKeys {
			log.Printf("Fetching videos with API key %s and query %s", apiKey, query)

			videos, err := FetchVideos(apiKey, query)
			if err != nil {
				log.Printf("Error fetching videos with API key %s and query %s: %v", apiKey, query, err)
				continue
			}

			StoreVideos(db, videos)
		}

		time.Sleep(interval)
	}
}
