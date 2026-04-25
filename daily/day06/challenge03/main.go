package main

import "fmt"

/*
	Challenge 3: Count Even Numbers
	nums = [1, 2, 3, 4, 6]
	Output : 3
*/

func main() {

	var n int

	fmt.Print("Enter a size: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var evenCount int

	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			evenCount++
		}
	}

	fmt.Println("Even Numbers are:", evenCount)
}
