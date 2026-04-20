package main

import "fmt"

/*
	Challenge 2: Struct Method
	Create method:
		- func (u User) isAdult() bool
	Return:
		- true if age >= 18
*/

type User struct {
	Name string
	Age  int
}

func (u User) isAdult() bool {
	return u.Age >= 18
}

func main() {
	user := User{
		Name: "test",
		Age:  23,
	}

	ok := user.isAdult()

	if ok {
		fmt.Println("Your are good to go!!")
	} else {
		fmt.Println("You have to wait bruhh")
	}

}
