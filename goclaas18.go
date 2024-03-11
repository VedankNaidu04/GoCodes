// //18th February package main

// code 1

// import "fmt"

// func main() {
// 	a := 1
// 	b := 2
// 	c := 3.5
// 	d := 4
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Printf("c is of type %T\n", c)
// }

// code 2

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b = 6, "Hello"
// 	c, d := 7, "World!"

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Printf("b is of type %T\n", b)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var (
// 		a int
// 		b int    = 1
// 		c string = "hello"
// 	)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(a, b, c) // Println is for new line ( linenew )
// 	fmt.Printf("%d\n", a)
// 	fmt.Printf("%d,%d,%s\n", a, b, c) // %d for integer and %s for string 
// }

// var can be usedd inside and outside the functions, variable declaration and value assignment can be done seperately
// := 
// 

package main 
import ("fmt")
 a=10; var b int =20
func main() {
	fmt.Println(a,b)
// RESULT = .\goclaas18.go:63:2: syntax error: non-declaration statement outside function body 
}
 
