package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

// A global map where all key-value data will be stored
var data = make(map[string]interface{})
var mu sync.Mutex // lock for safe access to data

const dataFile = "data.json"

// --- Load data from file ---
func loadData() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("No previous data found, starting fresh...")
		return
	}
	json.Unmarshal(file, &data)
}

// --- Save data to file ---
func saveData() {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(dataFile, bytes, 0644)
}

// --- PUT /objects ---
// Endpoint to save new data
func putHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	var body struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	data[body.Key] = body.Value
	saveData()
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

// --- GET /objects/{key} ---
// Endpoint to return data of a single key
func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/objects/"):]

	mu.Lock()
	value, exists := data[key]
	mu.Unlock()

	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(value)
}

// --- GET /objects ---
// Endpoint to return all keys and their values
func getAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	loadData()

	// /objects handles both GET and PUT
	http.HandleFunc("/objects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			putHandler(w, r)
		} else if r.Method == http.MethodGet {
			getAllHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Separate handler for /objects/{key}
	http.HandleFunc("/objects/", getHandler)

	fmt.Println(" Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
