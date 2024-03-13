// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	var dd int = 20
// 	var mm int = 02
// 	var yy int = 2024
// 	var str string
// 	str = fmt.Sprintf("%02-%02-%04", dd, mm, yy)
// 	io.WriteString(os.Stdout, str)
// }

//

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var str1 string
// 	var str2 string
// 	var a int
// 	var b float64
// 	fmt.Printf("Enter string:")
// 	fmt.Scan(&str1, &str2, &a, &b)
// 	fmt.Printf("Result: %s%s%d%f\n", str1, str2, a, b)
// 	// fmt.Printf("Result: %s\n", str2)
// 	// fmt.Printf("Result: %d\n", a)
// 	// fmt.Printf("Result: %f\n", b)
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 123
// 	var b uint = 0
// 	//assign the value of a to the b
// 	//b = a //ERROR : cannot use a (variable of type int) as uint value in assignment
// 	b = uint(a)
// 	//Printing the both values
// 	fmt.Printf("a = %d\nb = %d\n", a, b)
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = -123
	var b uint
	var c int = 12
	b = uint(a)
	fmt.Printf("a = %d\nb = %d\n", a, b)
	// Printing the absolute value of 'c'
	fmt.Printf("Absolute value of c: %f\n", math.Abs(float64(c)))         //higher value (b = 18446744073709551493) as b is unit
	fmt.Printf("Sqrt value of c: %f\n", float32(math.Sqrt(float64((c))))) //sqrt -- uses float64
}
