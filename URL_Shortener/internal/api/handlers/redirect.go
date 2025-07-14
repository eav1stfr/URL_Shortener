package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"urlshortener/internal/cache"
	"urlshortener/internal/sqlconnect"
	"urlshortener/utils"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	fmt.Println(shortUrl)
	longUrl, err := cache.CheckCacheForRedirect(fmt.Sprintf("%s:%s", os.Getenv("BASE_URL"), shortUrl))
	if err == nil {
		http.Redirect(w, r, longUrl, http.StatusFound)
		fmt.Println("AFTER CHECKING CACHE:", longUrl)
		return
	}
	fmt.Println("NO CACHE:", longUrl)
	id := utils.Decode(shortUrl)
	longUrl, err = sqlconnect.Decode(id)
	if err != nil {
		if appErr, ok := err.(*utils.AppErr); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		return
	}
	_ = cache.InsertShortToLongUrlCache(longUrl, fmt.Sprintf("%s:%s", os.Getenv("BASE_URL"), shortUrl))
	http.Redirect(w, r, longUrl, http.StatusFound)
	fmt.Println(longUrl)
}
