package main

import "fmt"

/*
	Challenge 1: Nested Struct
	Create:
		- User
		- Address
	Print:
		 - User lives in Delhi
*/

type Address struct {
	City    string
	HouseNo int
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func main() {

	user := User{
		Name: "test",
		Age:  33,
		Address: Address{
			City:    "Delhi",
			HouseNo: 101,
		},
	}

	fmt.Println("User lives in", user.Address.City)
}
