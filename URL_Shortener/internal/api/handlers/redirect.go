package handlers

import (
	"net/http"
	"strings"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := strings.TrimSuffix(r.URL.Path, "/")

	originalUrl, err :=
}
