package main

import (
	"database/sql"
	"fmt"
	"os"
)

func StartDatabase() (db *sql.DB, err error) {
	dbPath := fmt.Sprintf("%s/go-hex.db", os.TempDir())
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return
	}

	err = db.Ping()
	return
}
