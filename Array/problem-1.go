package main

import "fmt"

func main() {
	// Declaring an array of integers with a size of 5
	var arr [5]int
	fmt.Println("Empty array: ", arr)

	arr[0] = 10
	arr[2] = 56
	arr[4] = 100
	arr[3] = 90
	fmt.Println("Array after assignment: ", arr)

	// Declaring and innitializing an array at the same time
	languages := [4]string{"Go", "Python", "Java", "JavaScript"}
	fmt.Println("Array initialized with values: ", languages)

	fmt.Println("Looping using for loop ...")

	// looping the array using for loop iteration
	for i := 0; i < len(languages); i++ {
		fmt.Println("Element at index", i, "is", languages[i])
	}

	fmt.Println("Looping using Ranges...")

	// using Range
	for index, value := range languages {
		fmt.Println("Element at index", index, "is", value)
	}

	// accessing array elements using slice
	slice := languages[1:3]
	fmt.Println("Slice: ", slice)

	// updating array elements using slice
	slice[0] = "C++"
	fmt.Println("Updated slice: ", slice)
	fmt.Println("Updated array: ", languages)
}
