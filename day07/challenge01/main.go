package main

import "fmt"

/*
	Challenge 1: Map Practice
	Create a map:
		- name
		- age
		- city
	Print All values
*/

func main() {

	user := map[string]string{
		"name": "test",
		"age":  "12",
		"city": "delhi",
	}

	for a, b := range user {
		fmt.Println(a, ":", b)
	}
}
