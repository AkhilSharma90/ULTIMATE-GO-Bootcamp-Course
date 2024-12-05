package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	tasks   = []Task{}
	counter = 1
	mutex   = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		var task Task
		json.NewDecoder(r.Body).Decode(&task)
		mutex.Lock()
		task.ID = counter
		counter++
		tasks = append(tasks, task)
		mutex.Unlock()
		json.NewEncoder(w).Encode(task)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
