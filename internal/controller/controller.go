package controller

import (
	"encoding/json"
	"net/http"

	"github.com/priince9381/Url-Shortener/app/internal/database"
	"github.com/priince9381/Url-Shortener/app/internal/utils"
	"gorm.io/gorm"
)

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
			"short_url": url.ShortURL,
			"message":   "URL already exists",
		}
		sendJSONResponse(w, http.StatusOK, response)
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		sendJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: "Database error"})
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
		sendJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: "Failed to create short URL"})
		return
	}

	response := map[string]string{
		"short_url": shortURL,
	}
	sendJSONResponse(w, http.StatusOK, response)
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
		sendJSONResponse(w, http.StatusBadRequest, ErrorResponse{Message: "hort URL is incorrect"})
		return
	}
	sendJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: "Database error"})
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
