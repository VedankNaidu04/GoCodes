// 23 FEB
// *********************calculate progrm to calculate the sum of possitive numbers (AP)
package main

import "fmt"

func sum(number int) int {
	//condition to break recursion
	if number == 0 {
		return 0
	} else {
		return number + sum(number-1)
	}
}
func main() {

	var num int
	fmt.Printf("Enter number :")
	fmt.Scan(&num)
	if num < 0 {

		fmt.Printf("Invalid input")
	} else {
		// function call
		var result = sum(num)
		fmt.Println("Sum:", result)
	}

}
