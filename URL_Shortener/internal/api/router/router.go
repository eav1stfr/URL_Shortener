package router

import (
	"net/http"
	"urlshortener/internal/api/handlers"
)
import "github.com/go-chi/chi/v5"

func ConfigRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/ping", handlers.PingServer)

	return r
}
