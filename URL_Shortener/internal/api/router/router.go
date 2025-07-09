package router

import (
	"net/http"
	"urlshortener/internal/api/handlers"
)
import "github.com/go-chi/chi/v5"

func ConfigRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{shortUrl}", handlers.RedirectHandler)
	r.Post("/encode", handlers.EncodeUrl)

	return r
}
