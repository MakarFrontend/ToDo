package getDB

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const connToDB string = "user=postgres password=innerjoin dbname=ToDo sslmode=disable"

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", connToDB)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return db, nil
}
