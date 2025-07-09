package router

import (
	"github.com/go-chi/chi/v5"
	"urlshortener/internal/api/handlers"
)

func registerEncodePath(r chi.Router) {
	r.Route("/encode", func(r chi.Router) {
		r.Post("/", handlers.EncodeUrl)

	})
}
