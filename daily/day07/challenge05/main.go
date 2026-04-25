package main

import "fmt"

/*
	Challenge 5 - Monk
	Create function:
		- takes slice of students
		- returns student with highest marks
*/

type Student struct {
	Name  string
	Marks int
}

func getStudentMaxMarks(students []Student) Student {

	var finalStudent Student
	var maxMarks int = -1e9 // -1000000009

	for _, student := range students {
		if maxMarks < student.Marks {
			maxMarks = student.Marks
			finalStudent = student
		}
	}

	return finalStudent
}

func main() {

	var n int

	fmt.Print("Enter a size: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		var student Student
		fmt.Scan(&student.Name)
		fmt.Scan(&student.Marks)
		students = append(students, student)
	}

	finalStudent := getStudentMaxMarks(students)

	fmt.Println("The Student with Maximum marks is", finalStudent.Name, "and marks is", finalStudent.Marks)
}
