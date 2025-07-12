package sqlconnect

import "urlshortener/utils"

func IncrementClickCounter(id int64) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ConnectingToDatabaseError
	}
	defer db.Close()

	query := "UPDATE urls SET click_count = click_count + 1 WHERE id = $1"

	res, err := db.Exec(query, id)
	if err != nil {
		return utils.DatabaseQueryError
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return utils.DatabaseQueryError
	}
	if rows == 0 {
		return utils.UnitNotFoundError
	}

	return nil
}
