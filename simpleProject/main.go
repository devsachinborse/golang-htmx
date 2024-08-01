package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Book struct {
	Title  string
	Author string
}

var (
	books      []Book
	booksMutex sync.Mutex
	tmpl       = template.Must(template.ParseFiles("templates/index.html", "templates/book.html"))
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	router := mux.NewRouter()
	router.HandleFunc("/", formHandler).Methods("GET")
	router.HandleFunc("/submit", submitHandler).Methods("POST")

	fmt.Printf("server is running on http://localhost%s", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	booksMutex.Lock()
	defer booksMutex.Unlock()
	tmpl.ExecuteTemplate(w, "index.html", books)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		author := r.FormValue("author")

		book := Book{Title: title, Author: author}

		booksMutex.Lock()
		books = append(books, book)
		booksMutex.Unlock()

		tmpl.ExecuteTemplate(w, "book.html", books)
	}
}
