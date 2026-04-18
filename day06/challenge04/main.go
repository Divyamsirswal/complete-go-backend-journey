package main

import "fmt"

/*
	Challenge 4: Reverse Slice
	nums = [1, 2, 3, 4]
	Output = [4, 3, 2, 1]
*/

func main() {
	var n int

	fmt.Print("Enter a size: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Print("Before reversing slice: ")

	for i := 0; i < n; i++ {
		fmt.Print(nums[i], " ")
	}

	fmt.Println()

	// using two pointer approach here -> i guess it can handle both even and odd length
	s := 0
	e := n - 1

	for s <= e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}

	fmt.Print("After reversing slice: ")

	for i := 0; i < n; i++ {
		fmt.Print(nums[i], " ")
	}

}
