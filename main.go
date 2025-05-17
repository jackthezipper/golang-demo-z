package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

// Handler for POST requests
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received message from %s: %s", msg.Name, msg.Body)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s! Message received.", msg.Name)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "✅ Server is running and ready to receive files.")
}

func main() {
	http.HandleFunc("/message", postHandler)
	http.HandleFunc("/status", statusHandler) // <-- GET endpoint

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
