package router

import (
	"github.com/go-chi/chi/v5"
	"urlshortener/internal/api/handlers"
)

func registerRedirectPath(r chi.Router) {
	r.Route("/redirect", func(r chi.Router) {
		r.Post("/", handlers.RedirectHandler)
	})
}
