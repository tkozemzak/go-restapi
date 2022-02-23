package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book struct (model)
type Book struct {
	ID	string `json:"id"`
	Isbn	string `json:"isbn"`
	Title	string `json:"title"`
	Author	*Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname	string `json:"firstname"`
	Lastname	string `json:"lastname"`
}

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/api/books")
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//create new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}


func main() {
	//initialize router
	router := mux.NewRouter()

	//mock data

	//create route handlers/endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("Routes created")

	log.Fatal(http.ListenAndServe(":8000", router))
}