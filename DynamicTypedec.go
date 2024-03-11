// package main

// import "fmt"

// func main() {
// 	var x float64 = 20.0
// 	y := 242
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Printf("Value of x is %f\n", x)
// 	fmt.Printf("Value of y is %f\n", y)
// 	fmt.Printf("x is of type %T\n", x)
// 	fmt.Printf("y is of type %T\n", y)
// }

//Variable Declaration without Intial Value

// package main

// import "fmt"

// func main() {
// 	var a string
// 	var b int
// 	var c bool
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

//Mixed Variable Devlaration in Go

package main

import "fmt"

func main() {
	var a, b, c = 3, 4, "food"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("a is of type %t\n", a)
	fmt.Println("b is of type %t\n", b)
	fmt.Println("c is of type %t\n", c)
}
