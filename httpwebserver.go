package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserTemp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var usersTemp []UserTemp

func WebServer() {
	// Define endpoints
	http.HandleFunc("/users", handleUsers)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	case "POST":
		createUser(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Convert users to JSON
	jsonBytes, err := json.Marshal(usersTemp)
	if err != nil {
		http.Error(w, "Error converting users to JSON", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request body
	var newUser UserTemp
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Add user to list
	usersTemp = append(usersTemp, newUser)

	// Convert user to JSON
	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}
