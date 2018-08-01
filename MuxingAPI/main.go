package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
				)

// book struct (model)
type Book struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `'json: "title"'`
	Author *Author `'json: "author"'`
}

// author struct (model)
type Author struct {
	FirstName string `json: "firstname"`
	LastName string `json: "lastname"`
}

// init books var as sliceBook struct
var books []Book

// get all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	// loop through the books to find the ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update the book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// delete the book
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
	// this is the entry point
	// init the router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "48948489494", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Helmsworth"}})
	books = append(books, Book{ID: "2", Isbn: "54545455553", Title: "Book Two", Author: &Author{FirstName: "Allen", LastName: "Poems"}})
	books = append(books, Book{ID: "3", Isbn: "23333549494", Title: "Book Three", Author: &Author{FirstName: "Mike", LastName: "Dobison"}})
	books = append(books, Book{ID: "4", Isbn: "45645664591", Title: "Book Four", Author: &Author{FirstName: "Kevin", LastName: "Keating"}})
	books = append(books, Book{ID: "5", Isbn: "34453256777", Title: "Book Five", Author: &Author{FirstName: "Allison", LastName: "Bennington"}})
	books = append(books, Book{ID: "6", Isbn: "23444356879", Title: "Book Six", Author: &Author{FirstName: "George", LastName: "Smithers"}})

	// route handlers endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// listen and serve
	log.Fatal(http.ListenAndServe(":8080", r))
}
