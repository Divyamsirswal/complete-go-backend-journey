package main

import "fmt"

/*
	Challenge 2: Check Key Exists
	Create a map and Check if "email" exists in map
*/

func main() {

	user := map[string]string{
		"username": "test",
		"email":    "test@mail.com",
		"password": "xxxxxx",
	}

	value, ok := user["email"]
	if ok {
		fmt.Println("Value exists:", value)
	} else {
		fmt.Println("Not found")
	}
}
