package main

import (
	//"encoding/json"
	"log"
	//"math/rand"
	"net/http"
	//"strconv"

	"github.com/gorilla/mux"
)
// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
//init books var as a slice book struct
var books []Book
// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
}
// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
}
// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
}
// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
}
// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
}

func main() {
	// Init router
	r := mux.NewRouter()

	//Mock data -implement DB

	books = append (books,Book{ID: "1",Isbn:"787822", Title:"Book one",Author:&Author
	{Firstname:"chirag",Lastname:"painter"}})
	books = append (books,Book{ID: "2",Isbn:"745822", Title:"Book two",Author:&Author
	{Firstname:"smit",Lastname:"patel"}})

// Route handles & endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}	