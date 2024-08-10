package handlers

import (
	"encoding/json"
	"net/http"

	"20MIS0158_Backend/internal/models"
	"gorm.io/gorm"
)

// GetVideos returns the list of videos in a paginated response, sorted by published datetime in descending order.
func GetVideos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var videos []models.Video
		result := db.Order("published_at desc").Find(&videos)
		if result.Error != nil {
			http.Error(w, "Failed to retrieve videos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videos)
	}
}

// SearchVideos searches videos by title and description.
func SearchVideos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
			return
		}

		var videos []models.Video
		searchPattern := "%" + query + "%"
		result := db.Where("title LIKE ? OR description LIKE ?", searchPattern, searchPattern).Find(&videos)
		if result.Error != nil {
			http.Error(w, "Failed to search videos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videos)
	}
}
