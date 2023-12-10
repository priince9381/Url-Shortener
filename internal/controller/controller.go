package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"short_url": fmt.Sprintf("/%s", "Hello World"),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectToLongURL(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "www.google.com", http.StatusTemporaryRedirect)
}
