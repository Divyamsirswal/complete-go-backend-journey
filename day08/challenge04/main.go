package main

import "fmt"

/*
	Challenge 4: Slice of Nested Structs
	Create:
		- multiple users with addresses
		- loop and print all
*/

type Address struct {
	City  string
	State string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	var totalUsers int

	fmt.Print("Enter the number of users: ")
	fmt.Scan(&totalUsers)

	users := make([]User, totalUsers)

	for i := 0; i < totalUsers; i++ {
		var user User
		fmt.Print("Enter name: ")
		fmt.Scan(&user.Name)
		fmt.Print("Enter Age: ")
		fmt.Scan(&user.Age)
		fmt.Print("Enter city: ")
		fmt.Scan(&user.Address.City)
		fmt.Print("Enter state: ")
		fmt.Scan(&user.Address.State)

		users = append(users, user)
	}

	for _, val := range users {
		fmt.Println("Name:", val.Name)
		fmt.Println("Age:", val.Age)
		fmt.Println("City:", val.Address.City)
		fmt.Println("State:", val.Address.State)
	}

}
