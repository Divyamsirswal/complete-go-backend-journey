package main

import "fmt"

/*
	Challenge 1: Reverse String
	Input : aman
	Output : nama
*/

func main() {

	var name string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	var new_name string

	fmt.Print("Reversed Name: ")
	for i := len(name) - 1; i >= 0; i-- {
		new_name += string(name[i])
		fmt.Print(string(name[i]))
	}
	fmt.Println()
	fmt.Println("Final name after reversed: ", new_name)
}
