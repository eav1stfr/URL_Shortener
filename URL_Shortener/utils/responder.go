package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithShortUrl(w http.ResponseWriter, url string) {
	response := struct {
		Status string `json:"status"`
		Url    string `json:"short_url"`
	}{
		Status: "success",
		Url:    url,
	}
	json.NewEncoder(w).Encode(response)
}
