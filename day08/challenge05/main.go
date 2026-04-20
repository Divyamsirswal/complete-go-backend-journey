package main

import (
	"fmt"
)

/*
	Challenge 5 - Monk
	Create function:
		- takes slice of users
		- returns user from a specific city
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

func giveMeUserFromCity(users []User, city string) (User, error) {

	for _, user := range users {
		if user.Address.City == city {
			return user, nil
		}
	}
	fmt.Println("No user is present with that city !! :(")
	return User{}, fmt.Errorf("No user present")
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

	var city string

	fmt.Print("Enter a city from where you want the user: ")
	fmt.Scan(&city)

	newUser, err := giveMeUserFromCity(users, city)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newUser)
}
