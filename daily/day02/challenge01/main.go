package main

import "fmt"

/*
	Challenge 1: Personal Info System
	Take input:
		- name
		- age
		- city

	Output:
		- Hello Aman, you are 21 years old and live in Delhi.
*/

func main() {
	var name string
	var age int
	var city string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter your city: ")
	fmt.Scan(&city)

	fmt.Println("Information")
	fmt.Printf("Hello %s, you are %d years old and live in %s.", name, age, city)
}
