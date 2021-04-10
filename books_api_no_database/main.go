package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var booksDB []Book

var book Book

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/books", getbooks).Methods("GET")
	router.HandleFunc("/books/{id}", getbook).Methods("GET")
	router.HandleFunc("/books", addbook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	http.ListenAndServe(":8005", router)

}

func getbooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(booksDB)

	fmt.Println("GET all book is called and ")
}
func getbook(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	log.Println(reflect.TypeOf(val))
	id, _ := strconv.Atoi(val["id"])
	_, book := searchbook(id)
	if (book != Book{}) {
		json.NewEncoder(w).Encode(&book)
		fmt.Println("GET A BOOK IS CALLED")
	}
}

func searchbook(id int) (int, Book) {

	for loc, v := range booksDB {
		if id == v.ID {
			return loc, v
		}
	}
	return 0, Book{}
}
func addbook(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println("ERROR OCCURED")
	}
	booksDB = append(booksDB, book)
	fmt.Printf("ADD BOOK IS CALLED AND ADDED:  %v", booksDB)
	json.NewEncoder(w).Encode(booksDB)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	loc, bookval := searchbook(book.ID)
	if (bookval != Book{}) {
		fmt.Println(bookval.ID)
		booksDB[loc] = book
	} else {

		fmt.Fprintln(w, http.StatusNotFound)
	}

	fmt.Println("UPDATE BOOK IS CALLED")
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	id, _ := strconv.Atoi(value["id"])
	loc, val := searchbook(id) //searches for the book if present or not.
	if (val != Book{}) {
		booksDB = append(booksDB[:loc], booksDB[loc+1:]...)
	}
	json.NewEncoder(w).Encode(booksDB)
	fmt.Println("REMOVE BOOK IS CALLED")
}

