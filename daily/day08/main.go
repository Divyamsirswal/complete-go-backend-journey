package main

import "fmt"

type Address struct {
	City  string
	State string
}

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func (u User) greet() {
	fmt.Println("Hello", u.Name)
}

func updateAge(u *User) {
	u.Age = 30
}

func main() {

	user := User{
		Name: "aman",
		Age:  21,
		Address: Address{
			City:  "xyz",
			State: "test",
		},
	}

	user.greet()
	// fmt.Println(user.Address.City)

}
