package main

import "fmt"

// Defining a struct
type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s Student) getStudentDetails() {
	fmt.Printf("Name: %s, Age: %d, Grade: %d\n", s.Name, s.Age, s.Grade)
}

func main() {
	println("Struct in Go Lang which is equivalent to Objects")

	// Creating a new student object
	student := Student{
		Name:  "John Doe",
		Age:   18,
		Grade: 12,
	}

	// Calling the getStudentDetails method
	student.getStudentDetails()

	// Changing the student's grade
	student.Grade = 13

	// Calling the getStudentDetails method again to see the updated grade
	student.getStudentDetails()

}
