package database

import (
	"database/sql"
	"log"
)

func GetDatabase() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
