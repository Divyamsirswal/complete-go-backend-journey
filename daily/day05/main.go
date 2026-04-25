package main

import "fmt"

func basicFunction() {
	fmt.Println("Hello bruh")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {

	// basic function
	// basicFunction()

	// function with parameters
	// greet("Aman")

	// function with return value
	// res := add(2, 3)
	// fmt.Println(res)

	// function with multiple return value
	// divi, mod := divide(3, 4)
	// fmt.Println(divi, mod)
}
