package main

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "You hit GET Method")
	}

	w.WriteHeader(http.StatusOK)

	name := r.URL.Query()

	for k, v := range name {
		fmt.Print(k, "-")
		for _, str := range v {
			fmt.Print(str, " ")
		}
		fmt.Println()
	}

	path := r.URL.Path
	fmt.Println(path)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
