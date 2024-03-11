// package main

// import "fmt"

// func main() {
// 	var i, j string = "Hello", "World"

// 	fmt.Print(i)
// 	fmt.Print(j)
// }

// package main

// import "fmt"

// func main() {
// 	var Name, Branch, College_name string = "Vedank Naidu", "Cyber security", "RCOEM"
// 	var age int = 21

// 	fmt.Println("Name -", Name)
// 	fmt.Println("age -", age)
// 	fmt.Println("Branch -", Branch)
// 	fmt.Println("College Name -", College_name)

// }

// package main

// import "fmt"

// func main() {
// 	var a string = "Rajesh"
// 	fmt.Printf("value of a is'%s'\n", a)
// 	fmt.Println("value of a is", "'", a, "'")
// }

// package main
// import("fmt")

// func main(){
// 	const name, dept - "GFG", "CYSE"

// 	fmt.Printf("%s is a %s portal.\n", name, dept)
// }

// next code -11
// Golang program to illustrate the usage of
// fit. Printf() function
// Including the main package
package main

// Importing fat
import (
	"fmt"
)

// Calling main
func main() {
	var str = "Geeksforgeeks"
	fmt.Printf("The string is %s \n", str)
	var num1 int = 21
	fmt.Printf("The decimal value is %d \n", num1)
	var num2 float32 = 7.786
	fmt.Printf("The floating point is %g \n", num2)
	var num3 int = 14
	fmt.Printf("The binary value of num is %b \n", num3)
	var num4 = 4 + 1i
	fmt.Printf("Scientific Notation of numA : %e \n", num4)
	fmt.Printf("%v\n%v\n%v\n%v\n", str, num2, num3, num4) // use of %v is that use accordingly for it to interpret
}
