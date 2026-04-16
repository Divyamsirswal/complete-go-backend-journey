package main

import (
	"fmt"
	"unicode"
)

/*
	Challenge 2: Count Vowels
	Input : hello
	Output : 2
*/

func main() {

	var name string

	fmt.Print("Enter the string: ")
	fmt.Scan(&name)

	var vowels_count int

	for i := 0; i < len(name); i++ {
		char := unicode.ToLower(rune(name[i]))
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowels_count++
		}
	}

	fmt.Println("Total Vowels are: ", vowels_count)
}
