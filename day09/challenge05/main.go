package main

import "fmt"

/*
	Challenge 5 - Monk
	Create:
		- slice of users
	Function:
		- increase age of all users using pointer
*/

type User struct {
	Name string
	Age  int
}

func updateAllUsersAge(users []User) {

	for i := 0; i < len(users); i++ {
		var user *User
		user = &users[i]
		user.Age += 100
	}
}

func main() {
	users := []User{
		{Name: "u1", Age: 11},
		{Name: "u2", Age: 12},
		{Name: "u3", Age: 13},
		{Name: "u4", Age: 14},
		{Name: "u5", Age: 15},
	}

	updateAllUsersAge(users)

	fmt.Println("After updating age of all users")
	for _, user := range users {
		fmt.Println(user.Name, "->", user.Age)
	}
}
