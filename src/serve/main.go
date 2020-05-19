// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	db, err := sql.Open("postgres", "user=macosx password=123 dbname=testdb sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	if err != nil {
// 		// handle this error better than this
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		err = rows.Scan(&id, &name)
// 		if err != nil {
// 			// handle this error
// 			panic(err)
// 		}
// 		fmt.Println(id, name)
// 	}
// 	// get any error encountered during iteration
// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}
// }
package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Hello world::\n")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=macosx dbname=testdb password=123")
	defer db.Close()
	if err != nil {
		fmt.Println("Fail")
	} else {
		fmt.Println("Success")
	}
}