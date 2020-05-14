package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"Isbn"`
	Title  string  `json:"Title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

//init book
var books []Book

//function
func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "HomePage EndPoint Hit")
}

//get all book
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(w).EnCode(books)
}

//get single book by id
func getBook(w http.ResponseWriter, r *http.Request) {

}

//create book in the table
func createBook(w http.ResponseWriter, r *http.Request) {

}

//update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	//create router
	myRoute := mux.NewRouter().StrictSlash(true)

	//Mock data
	books = append(books, Book{
		ID:    "1",
		Isbn:   "44873",
		Title: "Tiếng gọi nơi hoang dã",
		Author: &Author{
			FistName: "John Biden",
			LastName: "Bidonn",
			Age:      45,
			Address:  "46 Khâm Thiên",
			Phone:    "0839828648",
			Email:    "banhbao01@mail.com",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:   "44321",
		Title: "1977 Vlogs",
		Author: &Author{FistName: "Chí Biden",
			LastName: "Phèo",
			Age:      40,
			Address:  "Làng Vũ Đại",
			Phone:    "0937275762",
			Email:    "chidepzai01@mail.com"
		}
	})

	//endpoint router
	myRoute.HandleFunc("/", homePage)
	myRoute.HandleFunc("/api/books", getBooks).Methods("GET")
	myRoute.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	myRoute.HandleFunc("/api/books", createBook).Methods("POST")
	myRoute.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	myRoute.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", myRoute))
}

func main() {
	handleRequests()
}
