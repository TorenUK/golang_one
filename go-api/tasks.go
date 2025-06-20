package main

import ( 
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"strconv"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var tasks = []Task{
	{ID: 1, Name: "Learn Go"},
	{ID: 2, Name: "Build an API"},
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
    idParam := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    for _, task := range tasks {
        if task.ID == id {
            json.NewEncoder(w).Encode(task)
            return
        }
    }

    http.Error(w, "Task not found", http.StatusNotFound)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	fmt.Printf("Recieved task: %+v\n", newTask)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", getAllTasks)
		r.Post("/", createTask)
		r.Get("/{id}", getTaskByID)
	})

	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}