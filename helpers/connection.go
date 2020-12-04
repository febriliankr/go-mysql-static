package helpers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("golang_app", "root:@tcp(127.0.0.1:3306)/database")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
