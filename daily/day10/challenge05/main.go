package main

import (
	"errors"
	"fmt"
)

/*
	Challenge 5 - Monk
	Create:
		- slice of numbers
	Function:
		- returns error if slice is empty
		- else return sum
*/

func calculateSum(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum, nil
}

func main() {
	var n int

	fmt.Print("Enter a size: ")
	fmt.Scanln(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	sum, err := calculateSum(nums)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", sum)
	}
}
