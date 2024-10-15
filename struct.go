package main

import "fmt"

// Defining a struct
type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) Greet() {
	fmt.Printf("Hello, I am %s. \n Currently I am %d years old \n", p.name, p.age)
}

func main() {
	// Creating a new instance of the Person struct
	person := Person{
		name:   "Tanvirul Islam",
		age:    25,
		gender: "male",
	}
	person.Greet()
}
