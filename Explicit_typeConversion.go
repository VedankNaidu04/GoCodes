package main

import "fmt"

func main() {
	var x int = 42
	var y float64 = float64(x)
	var z uint = uint(y)
	// printing the types and values of the variables
	fmt.Printf("Values of x is %d and type is %T\n", x, x)
	fmt.Printf("Values of y is %.2f and type is %T\n", y, y)
	fmt.Printf("Values of z is %d and type is %T\n", z, z)
}
