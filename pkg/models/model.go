package models

import (
	"database/sql"
	"log"
)

type Model interface {
	Insert (db *sql.DB, params ...interface{}) error
}

func GetID(db *sql.DB, selectQuery,insertQuery string, args ...interface{}) int {
	var id int
	err := db.QueryRow(selectQuery, args).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}