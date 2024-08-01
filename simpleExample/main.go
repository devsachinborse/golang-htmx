package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	port := ":3000"
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/time", handleTime)

	fmt.Printf("server running on http://localhost%s", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "index.html")
	teml, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Message string
	}{
		Message: "Hello world",
	}

	teml.Execute(w, data)
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Current time: " + time.Now().Format(time.RFC1123)))
}
