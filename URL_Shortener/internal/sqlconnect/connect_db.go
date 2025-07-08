package sqlconnect

import (
	"github.com/jmoiron/sqlx"
	"os"
	"urlshortener/utils"
)

func ConnectDb() (*sqlx.DB, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	if err = db.Ping(); err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	return db, nil
}
