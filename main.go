package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
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

//init books var as a slice Book struct
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params

	//loop through books and find correct id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) //mock id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...)
		break	
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {
	//initialize router
	router := mux.NewRouter()

	//mock data - todo: implement DB
	books = append(books, Book{ID: "1", Isbn: "243235", Title: "Book 1", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "343235", Title: "Book 2", Author: &Author{Firstname: "Bill", Lastname: "Murray"}})
	books = append(books, Book{ID: "3", Isbn: "543235", Title: "Book 3", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})


	//create route handlers/endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("Routes created")

	port := ":8000"

	fmt.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}