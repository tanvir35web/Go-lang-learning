package main

import "fmt"

func main() {
	fmt.Println("Hello, I am Variable ..")

	var name = "Tanvir"
	fmt.Println("My name is ", name)

	var age int = 25
	fmt.Println("I am ", age, " years old")

	var isMarried bool = true
	fmt.Println("Am I married? ", isMarried)

	var city string = "Dhaka"
	fmt.Println("I live in ", city)

	var b, c int = 2, 6
	fmt.Println("Sum of ", b, " and ", c, " is ", b+c)

	var d float64 = 3.14159
	fmt.Println("Pi value is approximately ", d)

	var e string = `This is a multi-line string.
	It can contain multiple lines and escape characters.`
	fmt.Println(e)

	Brand := "Apple"
	fmt.Println("I love ", Brand)

	number := 170
	fmt.Println("Square of ", number, " is ", number*number)

}
