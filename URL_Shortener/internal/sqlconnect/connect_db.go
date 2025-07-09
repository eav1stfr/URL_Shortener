package sqlconnect

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"urlshortener/utils"
)

func ConnectDb() (*sqlx.DB, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Println("CONNECTION STRING IS", connectionString)
		log.Println("ERR HERE 1")
		return nil, utils.ConnectingToDatabaseError
	}
	if err = db.Ping(); err != nil {
		return nil, utils.ConnectingToDatabaseError
		log.Println("ERR HERE 2")
	}
	return db, nil
}
