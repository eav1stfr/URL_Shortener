package sqlconnect

import (
	"database/sql"
	"log"
	"urlshortener/utils"
)

func Decode(id int64) (string, error) {
	db, err := ConnectDb()
	if err != nil {
		return "", err
	}
	defer db.Close()

	query := "SELECT long_url FROM urls WHERE id = $1"
	var url string
	err = db.Get(&url, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", utils.UnitNotFoundError
		}
		log.Println("ERR HERE 2")
		return "", utils.DatabaseQueryError
	}

	return url, nil
}
