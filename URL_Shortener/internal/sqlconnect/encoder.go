package sqlconnect

import (
	"database/sql"
	"time"
	"urlshortener/internal/cache"
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
	err = IncrementClickCounter(int64(id))
	return id, nil
}

func CheckExistence(url string) (string, error) {
	db, err := ConnectDb()
	if err != nil {
		return "", utils.ConnectingToDatabaseError
	}
	defer db.Close()
	query := "SELECT short_url, id FROM urls WHERE long_url = $1"
	var shortUrl string
	var id int64
	err = db.QueryRow(query, url).Scan(&shortUrl, &id)
	if err == sql.ErrNoRows {
		return "", utils.UnitNotFoundError
	} else if err != nil {
		return "", utils.DatabaseQueryError
	}
	err = IncrementClickCounter(id)
	if err != nil {
		return "", err
	}
	_ = cache.InsertLongToShortUrlCache(url, shortUrl)
	return shortUrl, nil
}

func AddShortUrl(shortUrl, longUrl string) error {
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE urls SET short_url = $1 WHERE long_url = $2"
	res, err := db.Exec(query, shortUrl, longUrl)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.DatabaseQueryError
	}
	if rowsAffected == 0 {
		return utils.DatabaseQueryError
	}
	_ = cache.InsertLongToShortUrlCache(longUrl, shortUrl)
	return nil
}
