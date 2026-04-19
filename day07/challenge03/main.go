package main

import "fmt"

/*
	Challenge 3: Struct Practice
	Create struct:
		type Student struct {
		    Name string
		    Marks int
		}
	Create 2 students and print them
*/

type Student struct {
	Name  string
	Marks int
}

func main() {

	student1 := Student{
		Name:  "s1",
		Marks: 43,
	}
	student2 := Student{
		Name:  "s2",
		Marks: 68,
	}

	fmt.Println("Student 1:", student1)
	fmt.Println("Student 2:", student2)
}
