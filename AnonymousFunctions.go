// //GO Anonymous Function

// package main

// import "fmt"

// var anony = func() {
// 	fmt.Println("HIHIHIHI")

// }

// func main() {
// 	// assigning anonymous function to variable
// 	var greet = func() {
// 		fmt.Println("Hello doston")
// 	}
// 	// function call
// 	var welcome = anony
// 	greet()
// 	welcome()
// }

//*************** Anonymous functions with paratmeters

// package main

// import "fmt"

// func main() {
// 	//anonymous functions with arguments
// 	var sum = func(n1 int, n2 int) {
// 		result := n1 + n2
// 		fmt.Printf("the sum of n1 and n2 is %d", result)
// 		//return result
// 	}
// 	//function
// 	sum(3, 5)
// }

// *************** Result Value From Anonymous functions
// package main

// import "fmt"

// func main() {
// 	//anonymous functions
// 	var sum = func(n1, n2 int) int {
// 		sum := n1 + n2
// 		return sum
// 	}

// 	//function call
// 	result := sum(4, 5)
// 	fmt.Println("Sum is :", result)
// }

//******************* Anonymous functions as Arguments to other functions

package main

import "fmt"

var sum = 0

// regular function to calculate square of numbers
func findSquare(num int) int {
	square := num * num
	return square
}
func main() {
	// anonymous function that returns sum of numbers

	sum := func(number1 int, number2 int) int {
		return number1 + number2
	}
	// function call
	result := findSquare(sum(6, 9))
	fmt.Println("Result is:", result)
}
