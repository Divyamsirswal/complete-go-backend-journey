package main

import "fmt"

/*
	Challenge 3: Pointer Update
	Create function:
		- takes *User
		- updates age

	- Verify change reflects in original
*/

type User struct {
	Name string
	Age  int
}

func updateAge(u *User, age int) {
	u.Age = age
}

func main() {
	var age int

	fmt.Print("Enter the new age for the user: ")
	fmt.Scan(&age)

	user := User{
		Name: "test",
		Age:  12,
	}

	updateAge(&user, age)

	fmt.Println("User new updated age is", user.Age)

}
