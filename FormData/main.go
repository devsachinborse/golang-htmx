package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Name  string
	Email string
}

func main() {
	port := ":3000"

	router := mux.NewRouter()
	router.HandleFunc("/", formHandler).Methods("GET")
	router.HandleFunc("/submit", submitHandler).Methods("POST")

	fmt.Printf("server running on http://localhost%s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		data := Data{
			Name:  name,
			Email: email,
		}

		tmpl.ExecuteTemplate(w, "result.html", data)
	}
}
