// OPERATORS
// Operators are used to perform operations on variables and values.
// An operator is a symbol that tells the compiler to perform specific mathematical or logical manipulations.
// go language is rich in bullt-in operators and provides the following types of operators -
// • Arithmetic Operators
// • Relational Operators
// • Logical Operators
// • Bitwise Operators
// • Assignment Operators
// • Miscellaneous Operators

// // ARITHEMATIC OPERATORS
// package main

// import "fmt"

// func main() {
// 	var n1, n2, sum, diff, pro, div, mod, x int
// 	n1 = 356
// 	n2 = 113
// 	x = 10
// 	sum = n1 + n2
// 	diff = n1 - n2
// 	pro = n1 * n2
// 	div = n1 / n2
// 	mod = n1 % n2
// 	x++
// 	fmt.Println("sum is", sum)
// 	fmt.Println("difference is", diff)
// 	fmt.Println("product is", pro)
// 	fmt.Println("division is", div)
// 	fmt.Println("Remainder is ", mod)
// 	fmt.Println(x)
// 	// fmt.Println(x++) yeh nahi hota hai
// 	// fmt.Printf(x++)
// 	// fmt.Print(x++)
// 	// printf(x++)
// 	// println(x++)
// 	// print(x++)
// }

//ASSIGNMENT OPERATORS

// package main

// import "fmt"

// func main() {
// 	var a int
// 	var x = 10
// 	x += 5
// 	fmt.Println(x)

// 	fmt.Println("Enter a no. - ")
// 	fmt.Scanf("%d", a)
// }

// package main

// import "fmt"

// func main() {
// 	var c int
// 	var a int = 5
// 	b := 3
// 	c = a & b
// 	fmt.Printf("%d", c)
// 	c = a | b
// 	fmt.Printf("%d", c)
// 	c = a ^ b
// 	fmt.Printf("%d", c)
// }

// Assignment Operators Continue...
// Operator                 Example                   Same As
// =                         x = 5                    x= 5
// +=                        X+=3                     x= x+ 3
// -=                        x-= 3                    X=X-3
// *=                        x *= 3                   x = x* 3
// /=                        x/= 3                    x= x/ 3
// %=                        x %= 3                   x = x % 3
// &=                        X8=3                     x = x &3
// |=                        x|= 3                    x= x| 3
// ^=                        X^=3                     X=x^3
// ＞>＝ (right shift)       x >>= 3                  x = x >> 3
// ＜＜= (left shift)        x＜＜＝3                  X＝X<<3

//COMPARISON/RELATIONAL OPERATORS

// package main

// import "fmt"

// func main() {

// 	a := 10
// 	fmt.Println(a > 5)
// 	fmt.Println(a == 5)
// 	fmt.Println(a < 5)

// 	// a = 20
// 	// b := 5
// 	// fmt.Println(a && b)
// 	// fmt.Println(a || b)
// 	// fmt.Println(!a && b)
// 	a = 20
// 	b := 5
// 	fmt.Println(a > 5 && b > 5)    // Corrected logical AND operation
// 	fmt.Println(a > 5 || b > 5)    // Corrected logical OR operation
// 	fmt.Println(!(a > 5) && b > 5) // Corrected logical NOT operation
// }

// Sizeof() OPERATOR

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	var num1 int = 10
// 	var num2 byte = 20

// 	fmt.Printf("Size of num1: %d\n", unsafe.Sizeof(num1))
// 	fmt.Printf("Size of num2: %d", unsafe.Sizeof(num2))
// }

//Explain the function of size of ?????

//FIND THE VARIABLE TYPE USING reflect.Typeof() FUNCTION

// Golang program to find the variable type
// using reflect.TypeOF() function

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {

// 	// defining the variables
// 	a := 10
// 	b := 10.20
// 	c := "Hello"
// 	d := true
// 	//e := []string("India", "USA", "UK") // string array

// 	// printing the type of the variables

// 	fmt.Println("a =", a)
// 	fmt.Println("b =", b)
// 	fmt.Println("c =", c)
// 	fmt.Println("d =", d)
// 	//fmt.Println("e =", e)

// 	fmt.Println("Type of a", reflect.TypeOf(a))
// 	fmt.Println("Type of b =", reflect.TypeOf(b))
// 	fmt.Println("Type of c", reflect.TypeOf(c))
// 	fmt.Println("Type of d =", reflect.TypeOf(d))
// 	//fmt.Println("Type of e =", reflect.TypeOf(e))
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	a := 10
// 	b := 48.39
// 	c := "ballu"
// 	d := true
// 	e := []string{"india", "usa"}

// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	fmt.Println("c:", c)
// 	fmt.Println("d:", d)
// 	fmt.Println("e:", e)

// 	fmt.Println("type of a:", reflect.ValueOf(a).Kind())
// 	fmt.Println("type of b:", reflect.ValueOf(b).Kind())
// 	fmt.Println("type of c:", reflect.ValueOf(c).Kind())
// 	fmt.Println("type of d:", reflect.ValueOf(d).Kind())
// 	fmt.Println("type of e:", reflect.ValueOf(e).Kind())
// }

// package main

// import "fmt"

// func main() {
// 	defer fmt.Println("Hello")
// 	defer fmt.Println("Hiii")
// 	fmt.Println("gppd morning")
// }

// package main

// import "fmt"

// func main() {
// 	var radius float32 = 0
// 	var area float32 = 0
// 	const PI = 3.14
// 	fmt.Printf("Enter radius-")
// 	fmt.Scanf("%f", &radius)

// 	area = PI * radius * radius
// 	fmt.Printf("area of circle is :%f", area)
// }

// package main

// import "fmt"

// func main() {
// 	var radius float32 = 0
// 	var area float32 = 0
// 	var perimeter float32 = 0

// 	const PI = 3.14
// 	fmt.Printf("Enter radius-")
// 	fmt.Scanf("%f", &radius)

// 	area = PI * radius * radius
// 	fmt.Printf("area of circle is :%f", area)
// 	fmt.Printf("Perimeter of Circle - ", perimeter)
// }



// GOLANG PROGRAM TO CONVERT THIS FAHRENHEIT TO CELCIUS 
package main

import "fmt"

func main() {
	var ftemp float64 = 0
	var ctemp float64 = 0
	fmt.Printf("Enter temperature in Fahrenheit -:")
	fmt.Scanf("%f", &ctemp)
	ctemp = (ftemp - 32) /1.8
	fmt.Printf("Temperature in Celsius: % 2f", ftemp)
}



// GOLANG PROGRAM TO CONVERT THIS CELCIUS TO FAHRENHEIT 
package main

import "fmt"

func main() {
	var ftemp float64 = 0
	var ctemp float64 = 0
	fmt.Printf("Enter temperature in celcius:")
	fmt.Scanf("%f", &ctemp)
	ftemp = (ctemp * 1.8) + 32
	fmt.Printf("Temperature in Fahrenheit : % 2f", ftemp)
}
