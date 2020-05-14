package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

	json.NewEncoder(w).Encode(books)
}

//get single book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params send by url
	//
	// loop
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create book in the table
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) ///mock id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] ///mock id
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params send by url
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func handleRequests() {
	//create router
	myRoute := mux.NewRouter().StrictSlash(true)

	//Mock data
	books = append(books, Book{
		ID:    "1",
		Isbn:  "44873",
		Title: "Tiếng gọi nơi hoang dã",
		Author: &Author{
			FirstName: "John Biden",
			LastName:  "Bidonn",
			Age:       45,
			Address:   "46 Khâm Thiên",
			Phone:     "0839828648",
			Email:     "banhbao01@mail.com",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "44321",
		Title: "1977 Vlogs",
		Author: &Author{
			FirstName: "Chí Biden",
			LastName:  "Phèo",
			Age:       40,
			Address:   "Làng Vũ Đại",
			Phone:     "0937275762",
			Email:     "chidepzai01@mail.com",
		},
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
