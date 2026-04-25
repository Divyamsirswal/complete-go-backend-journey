package main

type User struct {
	Name string
	Age  int
}

func updateAge(u *User) {
	u.Age = 23
}

func main() {
	// x := 10
	// p := &x

	// fmt.Println("address:", p)
	// fmt.Println("value:", *p)

	user := User{
		Name: "test",
		Age:  12,
	}

	updateAge(&user)

}
