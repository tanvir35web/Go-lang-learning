package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("Exam result ...")

	// for i := 0; i <= 20; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, " is even Number")
	// 	} else {
	// 		fmt.Println(i, " is odd Number")
	// 	}
	// }

	var examNumber int8 = 65

	if examNumber >= 80 && examNumber <= 100 {
		fmt.Println("You GOT A+ in the exam")
	} else if examNumber >= 60 && examNumber < 80 {
		fmt.Println("You GOT A in the exam")
	} else if examNumber >= 33 && examNumber < 60 {
		fmt.Println("You GOT B in the exam")
	} else if examNumber > 100 {
		fmt.Println("Invalid exam number")
	} else {
		fmt.Println("You failed the exam")
	}
}
