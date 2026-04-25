package main

import "fmt"

/*
	Challenge 1: Sum of Array
	nums := []int{1, 2, 3, 4, 5}
	Output: 15
*/

func main() {
	var n int

	fmt.Print("Enter a size: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	fmt.Println("Sum:", sum)

}
