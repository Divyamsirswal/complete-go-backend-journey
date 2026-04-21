package main

import "fmt"

/*
	Challenge 2: Update Value
	Function:
		- func change(n *int)
	change value to 100
*/

func change(x *int) {
	*x = 100
}

func main() {

	x := 19
	change(&x)

	fmt.Println(x)
}
