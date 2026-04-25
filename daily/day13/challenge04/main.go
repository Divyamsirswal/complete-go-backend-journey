package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 4: Status Codes
		- POST → 201 (created)
		- Error → 400 (bad request)
*/

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(map[string]string{
		"message": "working fine",
	}); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

}

func main() {
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
