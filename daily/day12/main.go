package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//JSON
/*
	{
		"name" : "test",
		"age" : 12,
	}
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
}

func main() {

	// user := User{
	// 	Name: "test",
	// 	Age:  12,
	// }

	// go struct ->  conversion -> JSON
	// data, _ := json.Marshal(user)

	// fmt.Println(string(data))

	http.HandleFunc("/", handler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
