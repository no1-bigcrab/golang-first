package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've requested the book:")
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've requested the book:")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've requested the book:")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've requested the book:")
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've requested the book:")
}

func router() {
	r := mux.NewRouter()

	//http Methods
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	//domain và host
	r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	//schemes router
	r.HandleFunc("/secure", BookHandler).Schemes("https")
	r.HandleFunc("/insecure", BookHandler).Schemes("http")

	//subrouter
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", BookHandler)
	bookrouter.HandleFunc("/{title}", BookHandler)

	// function cơ bản
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":8081", r)
}
