package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test title",
			Desc:    "test Description",
			Content: "Test Content",
		},
	}
	fmt.Println("Endpoint Hit: all article EndPoints")
	json.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage EndPoint Hit")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	myRoute := mux.NewRouter().StrictSlash(true)

	myRoute.HandleFunc("/", homePage)
	myRoute.HandleFunc("/allArticles", allArticles).Methods("GET")
	myRoute.HandleFunc("/allArticles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRoute))
}

func main() {
	handleRequests()
}
