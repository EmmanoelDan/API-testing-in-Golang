package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", GetBooks).Methods("GET")

	router.HandleFunc("/books", CreateBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
