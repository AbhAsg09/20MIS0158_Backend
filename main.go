package main

import (
	"20MIS0158_Backend/config"
	"20MIS0158_Backend/internal/db"
	"20MIS0158_Backend/internal/handlers"
	"20MIS0158_Backend/internal/models"
	"20MIS0158_Backend/internal/yt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the database with configuration
	db.InitDB(cfg)

	// Connect to the database
	gormDB := db.DB

	// Auto-Migrate Video model
	err = gormDB.AutoMigrate(&models.Video{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	// Initialize the router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/videos", handlers.GetVideos(gormDB)).Methods("GET")
	router.HandleFunc("/search", handlers.SearchVideos(gormDB)).Methods("GET")

	// Start background YouTube data fetching
	go yt.FetchAndStoreVideos(gormDB, cfg.YouTube.APIKeys, cfg.YouTube.SearchQuery, cfg.GetFetchInterval())

	// Start the server
	log.Printf("Starting server on %s...", cfg.Server.Port)
	if err := http.ListenAndServe(cfg.Server.Port, router); err != nil {
		log.Fatal(err)
	}
}
