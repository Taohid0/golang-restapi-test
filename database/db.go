package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB= ConnecToDB()

func ConnecToDB() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cricket_test")

	if err != nil {
		log.Fatal(err)

	}
	return db
}
