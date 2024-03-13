// $$$The Ivalues and the values in Go$$$ -
// There are two kinds of expressions in Go -
// •Ivalue - Expressions that refer to a memory location is called "Ivalue" expression. An Ivalue may appear as either the left-hand or right-hand side of an assignment.
// ervalue - The term value refers to a data value that is stored at some address in memory. An rvalue is an expression that cannot have a value assigned to it which means an value may appear on the right- but not left-hand side of an assignment.
// Variables are Ivalues and so may appear on the left-hand side of an assignment. Numeric literals are rvalues and so may not be assigned and can not appear on the left-hand side,
// The following statement is valid -
// X=20.0
// The following statement Is not valld. It would generate compile-time error -
// 10 = 20

// $$$Integer value literals$$$:
// There are four integer value literal forms, the decimal (base 10) form, the octal (base 8) form, the hex (base 16) form and the binary form (base 2).
// For example, the following four integer literals all denote
// 15 in decimal.0xF // the hex form (starts with a "0x" or "0X") OXF 017 // the octal-form (starts with a "0", "0o" or "00")
// 0017 0017 0b1111 // the binary form (starts with a "0b" or "0B")
// 0B1111 15 // the decimal form (starts without a "0")
// (Note: the binary form and the octal from starting with 0o or 00 are supported since Go 1.13.)

// The following program will print two true texts-

// package main

// import ("fmt")

// func main() {
// 	println(15 == 017) //true
// 	println(15 == 0xF) //true
// 	// Printf("aaaoo") // printf anf Printf dosent work
// }

// increment and decrement in GO

//package main

// //expected output- Binary of 15 is 1111
// import (
// 	"fmt"
// )

// func main() {
// 	var num3 int = 15
// 	fmt.Printf("The binary value of %d", num3, "is", " %b \n", num3, num3)
// }

// $$$Format Specifiers$$$ :
// It tells Golang how to format different data types. Some of the most commonly used specifiers are:
// • v - formats the value in a default format
// • d - formats decimal integers
// • g - formats the floating-point numbers
// • b - formats base 2 numbers
// • 0 - formats base 8 numbers
// • t - formats true or false values(Boolean)
// • s - formats string values

// package main

// import "fmt"

// func main() {
// 	//fmt.Printf(15 == 0xf)
// 	fmt.Printf("Expresssion 15==0xf is %T", 15 == 0xf) //RESULT _ Expresssion 15==0xf is bool
// }

// package main
// // Printing multiple values at Once
// import "fmt"
//     fname:= "bhai "
// 	lname:= "bhai !!!"
// func main(){

// 	fmt.Print("my name is", fname,"and my last name is", lname)
// }

// package main

// import "fmt"

// func main() {
// 	var a, b string = "Hello", "World"
// 	{

// 		fmt.Printf("%s\n", a)
// 		fmt.Printf("%s\n", b)
// 	}
// 	{
// 		fmt.Printf("Hello")
// 		fmt.Printf(" World!!!")
// 	}
// 	{
// 		fmt.Printf("%s\n%s", a, b)
// 		fmt.Printf("%s%s", a, b)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var i = 15.5
// 	var txt = "Hello World!"
// 	fmt.Printf("%v\n", i)
// 	fmt.Printf("%#v|n", i)
// 	fmt.Printf("%v%%\n", i)
// 	fmt.Printf("%T\n", i)
// 	fmt.Printf("%v\n", txt)
// 	fmt.Printf("%#v\n", txt)
// 	fmt.Printf("%T\n", txt)
// }

// package main

// import "fmt"

// func main() {
// 	var i = 15

// 	fmt.Printf("%b\n", i)
// 	fmt.Printf("%d\n", i)
// 	fmt.Printf("%+d\n", i)
// 	fmt.Printf("%o\n", i)
// 	// fmt.Printf("%0\n", i) // This line seems to be incorrect, please clarify what you intended here.
// 	fmt.Printf("%x\n", i)
// 	fmt.Printf("%X\n", i)
// 	fmt.Printf("%#x\n", i)
// 	fmt.Printf("%4d\n", i)
// 	fmt.Printf("%-4d\n", i)
// 	fmt.Printf("%04d\n", i)
// }

package main

import "fmt"

func main() {
	var txt = "Hello"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("%x\n", txt)
}
