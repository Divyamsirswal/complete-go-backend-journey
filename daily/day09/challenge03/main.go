package main

import "fmt"

/*
	Challenge 3: Struct Update
	Create:
		- User struct
	Function:
		- update name using pointer
*/

type User struct {
	Name string
	Age  int
}

func updateUserName(u *User) {
	u.Name = "naaahhhh"
}

func main() {
	user := User{
		Name: "wow",
		Age:  12,
	}

	updateUserName(&user)

	fmt.Println("After update the user name is:", user.Name)
}
