package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/priince938/app/internal/database"
	"github.com/priince938/app/internal/utils"
	"gorm.io/gorm"
)

type CreateShortURLRequest struct {
	LongURL string `json:"long_url"`
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// Get the GORM DB instance
	db := database.GetDB()

	var requestBody CreateShortURLRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	longURL := requestBody.LongURL
	if longURL == "" {
		http.Error(w, "Missing 'original_url' parameter", http.StatusBadRequest)
		return
	}

	// Check if the long URL already exists in the database
	var url database.URL
	result := db.Where("original_url = ?", longURL).First(&url)
	if result.Error == nil {
		response := map[string]string{
			"short_url": fmt.Sprintf("/%s", url.ShortURL),
			"Message":   "Url Already Exist",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// If the long URL doesn't exist in the database, generate a new short URL
	shortURL := utils.EncodeingId(longURL)

	// Create a new record in the database for the new short URL
	newURL := database.URL{
		OriginalURL: longURL,
		ShortURL:    shortURL,
	}

	if err := db.Create(&newURL).Error; err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"short_url": fmt.Sprintf("/%s", shortURL),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	// Get the GORM DB instance
	db := database.GetDB()

	shortURL := r.URL.Path[len("/get_url/"):] // Extract the short URL from the request path

	var url database.URL
	result := db.Where("short_url = ?", shortURL).First(&url)
	if result.Error == nil {
		// If the short URL exists in the database, redirect to the corresponding long URL
		http.Redirect(w, r, url.OriginalURL, http.StatusTemporaryRedirect)
		return
	} else if result.Error == gorm.ErrRecordNotFound {
		http.Error(w, "Short URL is incorrect", http.StatusBadRequest)
		return
	}

	http.Error(w, "Database error", http.StatusInternalServerError)
}
