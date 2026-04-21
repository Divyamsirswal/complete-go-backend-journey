package main

import "fmt"

/*
	Challenge 1: Basic Pointer
	Create variable:
		- x := 10
	Print:
		- value
		- address
		- value using pointer
*/

func main() {

	x := 10
	ptr := &x

	fmt.Println("Value:", x)
	fmt.Println("Address:", &x)
	fmt.Println("Value using pointer:", *ptr)

}
