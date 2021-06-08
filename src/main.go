// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book - Our struct for all books
type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", getBooks)
	myRouter.HandleFunc("/book", createNewBook).Methods("POST")
	myRouter.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", getBookById)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Books = []Book{
		Book{Id: "1", Title: "Book1", Author: "Author 1"},
		Book{Id: "2", Title: "Book 2", Author: "Author 2"},
		Book{Id: "3", Title: "Book 3", Author: "Author 3"},
	}
	handleRequests()
}
