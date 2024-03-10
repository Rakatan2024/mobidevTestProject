package database

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateDB(DBDriver string) *sql.DB {
	db, err := sql.Open(DBDriver, "sqlite3DB/database.db")
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
	}
	return db
}
