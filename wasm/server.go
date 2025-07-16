package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Todo struct {
	Text string `json:"text"`
}

var (
	todos []Todo
	mu    sync.Mutex
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	mu.Lock()
	todos = append(todos, t)
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			addTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 