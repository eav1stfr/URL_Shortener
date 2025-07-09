package sqlconnect

import (
	"database/sql"
	"time"
	"urlshortener/utils"
)

// first check if exists
func EncoderDbHandler(url string) (int, error) {
	db, err := ConnectDb()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	query := "INSERT INTO urls (long_url, created_at) VALUES ($1, $2) RETURNING id"
	var id int
	err = db.Get(&id, query, url, time.Now())
	if err != nil {
		return 0, utils.DatabaseQueryError
	}
	return id, nil
}

func CheckExistence(url string) (string, error) {
	db, err := ConnectDb()
	if err != nil {
		return "", utils.ConnectingToDatabaseError
	}
	defer db.Close()
	query := "SELECT short_url FROM urls WHERE long_url = $1"
	var shortUrl string
	err = db.Get(&shortUrl, query, url)
	if err == sql.ErrNoRows {
		return "", utils.UnitNotFoundError
	} else if err != nil {
		return "", utils.DatabaseQueryError
	}
	return shortUrl, nil
}
