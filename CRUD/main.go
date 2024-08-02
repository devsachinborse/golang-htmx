package main

import (
	"devsachinborse/htmx-crud/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/tasks", listTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/new", handleNewTask).Methods("GET")
	router.HandleFunc("/tasks", handleCreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}/edit", handleEditTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handleUpdateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}/delete", handleDeleteTask).Methods("POST")

	// fs := http.FileServer(http.Dir("./static/"))
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	fmt.Printf("server running on http://localhost%s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

var tasks = []models.Task{}
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func listTaskHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", tasks)
}

func handleNewTask(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "form.html", nil)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	task := models.Task{
		ID:    len(tasks) + 1,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func handleEditTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for _, task := range tasks {
		if task.ID == id {
			tmpl.ExecuteTemplate(w, "form.html", task)
			return
		}
	}
	http.NotFound(w, r)

}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = r.FormValue("title")
			break
		}
	}
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
