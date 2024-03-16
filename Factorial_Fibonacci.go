// Factorial of a No.

package main

import "fmt"

// Function to calculate factorial using recursion
func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	} else {
		// In Recursive case: factorial of n is n * factorial of (n-1)
		return n * factorial(n-1)

	}
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	// Calculate factorial
	result := factorial(num)
	fmt.Printf("Factorial of %d is: %d\n", num, result)
}
