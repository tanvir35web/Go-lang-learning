// package main

// import "fmt"

// func main() {
// 	// Declaring a 2x3 array
// 	var matrix [2][3]int

// 	// Assigning values
// 	matrix[0][0] = 1
// 	matrix[0][1] = 2
// 	matrix[0][2] = 3
// 	matrix[1][0] = 4
// 	matrix[1][1] = 5
// 	matrix[1][2] = 6

// 	// Iterating through a 2D array
// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[i]); j++ {
// 			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	// Create a 3x3 matrix
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Print matrix values
	fmt.Println("Matrix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	// Access specific element
	fmt.Println("Element at [2][1]:", matrix[2][1])

	// Change a value in the matrix
	matrix[1][1] = 99
	fmt.Println("Matrix after change:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}
