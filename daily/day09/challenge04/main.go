package main

import "fmt"

/*
	Challenge 4: Swap Numbers
		func swap(a *int, b *int)
		swap values
*/

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 1
	b := 2

	swap(&a, &b)

	fmt.Println("After swap a:", a)
	fmt.Println("After swap b:", b)
}
