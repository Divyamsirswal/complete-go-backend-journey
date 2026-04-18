package main

import "fmt"

/*
	Challenge 2: Find Max Element
	nums = [3, 7, 2, 9, 5]
	Output : 9
*/

func main() {
	var n int

	fmt.Print("Enter a size: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var maxElement int
	for i := 0; i < n; i++ {
		maxElement = max(maxElement, nums[i])
	}

	fmt.Println("Maximum element is:", maxElement)
}
