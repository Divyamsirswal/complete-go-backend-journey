package main

import "fmt"

/**
	Challenge 2:
  	 - Change values and observe output. (of Challenge 1)
*/

func main() {
	name := "Divyam Sirswal"
	city := "xyz"
	age := 11
	isStudent := true

	fmt.Println("Personal Details")
	fmt.Println("Name : ", name)
	fmt.Println("City : ", city)
	fmt.Println("Age : ", age)
	fmt.Println("Are you Student : ", isStudent)

	name = "Peter Rawat"
	age = 22

	fmt.Println()
	fmt.Println("Personal Details after modification")
	fmt.Println("Name : ", name)
	fmt.Println("City : ", city)
	fmt.Println("Age : ", age)
	fmt.Println("Are you Student : ", isStudent)

}
