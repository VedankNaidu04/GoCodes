//22 Feb

// DEFINING FUNCTIONS *************

// package main
// import "fmt"
// func addNumbers(){
// 	n1:=12
// 	n2:=2

// 	sum:=n1 + n2
// 	fmt.Println("Sum:",sum)

// 	func main(){
// 		//function call
// 		addNumbers(n1,n2)
// 	}
// }

// ************* ADDING 2 USER DEFINED NUMBERS
// package main

// import "fmt"

// //defines a function with 2 parameters
// func addNumbers(n1 int, n2 int, n3 int) {
// 	sum := n1 + n2 + n3
// 	fmt.Println("Sum:", sum)
// }
// func main() {
// 	var x, y, z int
// 	fmt.Println("Enter x :")
// 	fmt.Scan(&x)
// 	fmt.Println("Enter y :")
// 	fmt.Scan(&y)
// 	fmt.Println("Enter z :")
// 	fmt.Scan(&z)

// 	//pass parameters in function call
// 	addNumbers(x, y, z)
// }

// ************* RETURN VALUE FROM GO FUNCTION

// package main

// import "fmt"

// func addNumbers(n1 int, n2 int) int {
// 	sum := n1 + n2
// 	return sum
// }
// func main() {
// 	result := addNumbers(21, 14)
// 	fmt.Println("Sum:", result)
// }

// package main

// import "fmt"

// func addNumbers(n1 int, n2 int) int {
// 	sum := n1 + n2
// 	return sum

// 	//code after return statement
// 	// fmt.Println("after return statement") // ERROR : missing return
// 	// return statement should be the last statement except it is a recursssive function
// }
// func main() {
// 	result := addNumbers(21, 14)
// 	fmt.Println("Sum:", result)
// }

//************* Return Multiple Vaalues

// package main

// import "fmt"

// func Calculate(n1 int, n2 int) (int, int) {
// 	sum := n1 + n2
// 	difference := n1 - n2
// 	return sum, difference
// }
// func main() {

// 	var x, y int

// 	fmt.Printf("Enter x:")
// 	fmt.Scan(&x)
// 	fmt.Printf("Enter y:")
// 	fmt.Scan(&y)

// 	sum, difference := Calculate(x, y)
// 	fmt.Println("sum:", sum, "diff:", difference)
// }

// *******************Benefits of Using Functions
// • 1. Code Reusability
// Program to illustrate code reusability in function

package main

import "fmt"

// function definition
func getSquare(num int) {

	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)
}

// main function
func main() {
	//call function 3 times
	getSquare(3)
	getSquare(5)
	getSquare(10)
}
