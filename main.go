package main

import (
	"20MIS0158_Backend/config"
	"20MIS0158_Backend/internal/db"
	"20MIS0158_Backend/internal/handlers" // Import your API handlers
	"20MIS0158_Backend/internal/yt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func startFetching(cfg *config.Config, gormDB *gorm.DB) {
	apiKeys := cfg.YouTube.APIKeys
	query := cfg.YouTube.SearchQuery

	log.Println("Starting video fetcher...") // Add logging

	yt.FetchAndStoreVideos(gormDB, apiKeys, query, cfg.GetFetchInterval())
}

func main() {
	// Load configuration

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the database with configuration
	db.InitDB(cfg)
	gormDB := db.DB

	// Auto-migrate the Video model
	if err := gormDB.AutoMigrate(&yt.Video{}); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	// Start background YouTube data fetching
	go startFetching(cfg, gormDB)

	// Setup routes and start the HTTP server
	http.HandleFunc("/videos", handlers.GetVideos(gormDB))    // Handle /videos endpoint
	http.HandleFunc("/search", handlers.SearchVideos(gormDB)) // Handle /search endpoint

	// Start the HTTP server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
