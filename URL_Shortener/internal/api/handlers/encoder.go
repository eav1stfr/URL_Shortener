package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"urlshortener/internal/cache"
	"urlshortener/internal/sqlconnect"
	"urlshortener/utils"
)

func EncodeUrl(w http.ResponseWriter, r *http.Request) {
	var longUrl string
	err := json.NewDecoder(r.Body).Decode(&longUrl)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayload.Error(), utils.InvalidRequestPayload.GetStatusCode())
		return
	}
	if !strings.HasPrefix(longUrl, "http") {
		http.Error(w, utils.InvalidRequestPayload.Error(), utils.InvalidRequestPayload.GetStatusCode())
		return
	}

	shortUrl, err := cache.CheckCacheForEncoding(longUrl)
	if err == nil {
		utils.RespondWithShortUrl(w, shortUrl)
		return
	}
	shortUrl, err = sqlconnect.CheckExistence(longUrl)
	if err == nil {
		utils.RespondWithShortUrl(w, shortUrl)
		return
	}
	id, err := sqlconnect.EncoderDbHandler(longUrl)
	if err != nil {
		if appErr, ok := err.(*utils.AppErr); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	code := utils.Encode(id)
	baseUrl := os.Getenv("BASE_URL")
	response := struct {
		Status string `json:"status"`
		Url    string `json:"url"`
	}{
		Status: "success",
		Url:    fmt.Sprintf("%s/%s", baseUrl, code),
	}
	err = sqlconnect.AddShortUrl(fmt.Sprintf("%s/%s", baseUrl, code), longUrl)
	if err != nil {
		if appErr, ok := err.(*utils.AppErr); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingMessageError.Error(), utils.EncodingMessageError.GetStatusCode())
	}
}
