package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Result Struct
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    []Book `json:"data,omitempty"`
}

// Book Struct (Model)
type Book struct {
	ID     int     `json:"id,omitempty"`
	Isbn   string  `json:"isbn,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author *Author `json:"author,omitempty"`
}

// Author Struct
type Author struct {
	Name string `json:"name,omitempty"`
	City string `json:"city,omitempty"`
}

// Init Books slice
var books []Book
var result Result

// Get books w 返回要写的对象 r 客户端请求的对象
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := Result{Message: "success", Data: books}
	json.NewEncoder(w).Encode(result)
}

// Get Single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if strconv.Itoa(item.ID) == params["id"] {
			var data []Book
			data = append(data, item)
			result = Result{Message: "success", Data: data}
			json.NewEncoder(w).Encode(result)
			return // 提前返回
		}
	}

	// 如果没有找到相应id的书,返回未找到
	result = Result{Code: 1, Message: "don't fund record"}
	json.NewEncoder(w).Encode(result)

}

// create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 直接得到json
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = rand.Intn(10000)

	// 插入数据
	books = append(books, book)

	result = Result{Message: "success"}
	json.NewEncoder(w).Encode(result)

}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i, item := range books {
		if strconv.Itoa(item.ID) == params["id"] {
			books = append(books[:i], books[i+1:]...)

			// 直接得到json
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			break
		}
	}

	result = Result{Message: "success"}
	json.NewEncoder(w).Encode(result)
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i, item := range books {
		if strconv.Itoa(item.ID) == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	result = Result{Message: "success"}
	json.NewEncoder(w).Encode(result)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    1,
		Isbn:  "18301928",
		Title: "a golang book",
		Author: &Author{
			Name: "zwhset",
			City: "beijing",
		},
	})
	books = append(books, Book{
		ID:    2,
		Isbn:  "q23123123",
		Title: "a python book",
		Author: &Author{
			Name: "l0set",
			City: "hunan",
		},
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Now Start Server...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
