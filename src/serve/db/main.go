package db

import (
	"fmt"
	pg "github.com/go-pg/pg"
)

func Connect(){
	opts := &pg.Options {
		User: "postgres",
		Password: "",
		Address: "localhost:5432",

	}

	var db *pg.DB = pg.Connect( opts)

	if db == nil {
		fmt.Printf("Connect data fail....\n")
		os.Exist(100)
	}
	fmt.Printf("Connect data success ....\n")

	closeErr:= db.Close()
	if closeErr == nil {
		fmt.Printf("Erro while closing the connect, data success: %v\n", closeErr)
		os.Exist(100)

	}
	fmt.Printf("Connect data success and closeing ....\n")
	return
}
