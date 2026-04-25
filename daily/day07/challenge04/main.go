package main

import "fmt"

/*
	Challenge 4: Slice of Structs
	Store 3 students in slice and loop through them
*/

type Student struct {
	Name  string
	Marks int
}

func main() {

	students := []Student{
		{
			Name:  "s1",
			Marks: 21,
		},
		{
			Name:  "s2",
			Marks: 45,
		},
		{
			Name:  "s3",
			Marks: 68,
		},
	}

	for _, student := range students {
		fmt.Println(student)
	}
}
