package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connectct bcbksjh
func Connect() {
	db, err := sql.Open("postgres", "user=macosx password=123 dbname=testdb sslmode=disable")
	if err != nil {
		fmt.Println("Fail")
		log.Fatal(err)
	} else {
		fmt.Println("Success")
	}
	defer db.Close()
	if err != nil {
		panic(err)
	}
}
