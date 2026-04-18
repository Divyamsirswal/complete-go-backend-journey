package main

import "fmt"

/*
	Challenge 5 - Monk

	Remove element from slice
	nums = [1, 2, 3, 4, 5] → remove index 2
	Output : [1, 2, 4, 5]
*/

func main() {
	var n int

	fmt.Print("Enter a size: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var index int

	fmt.Print("Enter the index where you want to remove element: ")
	fmt.Scan(&index)

	//lets check its under the bound
	if index < 0 && index > n-1 {
		fmt.Println("Please re-enter valid index")
		return
	}

	//one think i can do here by skipping that particular index i can use a new slice and append all other elements to it -> i can do this only for now

	nums2 := make([]int, 0)
	for i := 0; i < n; i++ {
		if i == index {
			continue
		}
		nums2 = append(nums2, nums[i])
	}

	fmt.Print("After Removing the element: ")
	for i := 0; i < len(nums2); i++ {
		fmt.Print(nums2[i], " ")
	}

	fmt.Println()
}
