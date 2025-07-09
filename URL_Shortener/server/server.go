package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"urlshortener/internal/api/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mux := router.ConfigRouter()
	server := &http.Server{
		Addr:    os.Getenv("API_PORT"),
		Handler: mux,
	}
	log.Println("server running on port 3000")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
