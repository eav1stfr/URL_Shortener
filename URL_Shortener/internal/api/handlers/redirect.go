package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"urlshortener/internal/sqlconnect"
	"urlshortener/utils"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside redirect")
	shortUrl := chi.URLParam(r, "shortUrl")
	id := utils.Decode(shortUrl)
	longUrl, err := sqlconnect.Decode(id)
	if err != nil {
		if appErr, ok := err.(*utils.AppErr); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		return
	}
	http.Redirect(w, r, longUrl, http.StatusFound)
	fmt.Println(longUrl)
}
