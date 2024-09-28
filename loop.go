package main

import "fmt"

func main() {
	// for loop practices

	// Example 1: Print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println("increment = ", i)
	}

	// Example 2: Print numbers from 10 to 1 in reverse order
	for i := 10; i >= 1; i-- {
		fmt.Println("decrement = ", i)
	}

	// Example 3: Print even numbers from 0 to 20
	i := 0
	for i <= 20 {
		fmt.Println("Even numbers = ", i)
		i += 2
	}

	// Example 4: Print odd numbers from 1 to 20
	for i := range 20 {
		if i%2 != 0 {
			fmt.Println("Odd numbers = ", i+1)
		}
	}

	// Example 5: Print numbers from 1 to 10, but only if they are prime numbers
	for i := 2; i <= 10; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println("Prime numbers = ", i)
		}
	}

	// Example 6: Print numbers from 1 to 10, but only if they are perfect squares
	for i := 1; i <= 10; i++ {
		isPerfectSquare := false
		for j := 1; j*j <= i; j++ {
			if j*j == i {
				isPerfectSquare = true
				break
			}
		}
		if isPerfectSquare {
			fmt.Println("Perfect squares = ", i)
		}
	}

}
