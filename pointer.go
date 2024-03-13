// 19FEB-pointer
// package main

// import "fmt"

// func main() {

// 	integer := 23 //variable is defined (as 23) w/o declaring it.
// 	fmt.Printf("%T%T\n", integer, &integer)
// 	fmt.Print("integer -", integer, &integer)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	number, floatingNumber := 238, 1234.575883939
// 	fmt.Printf("Deafault: %f\n", floatingNumber)
// 	fmt.Printf(".2f: %.2f\n", floatingNumber)
// 	fmt.Printf(".4f: %.4f\n", floatingNumber)
// 	fmt.Printf(".5f: %.5f\n", floatingNumber)
// 	fmt.Printf(".5.2f: %5.2f\n", floatingNumber)
// 	fmt.Printf("Scientific: %e\n", floatingNumber)
// 	fmt.Printf("Decimal: %d\n", number)
// 	fmt.Printf("Binary: %b\n", number)
// 	fmt.Printf("Octal: %o\n", number)
// 	fmt.Printf("HexaDecimal: %X\n", number)
// }

// package main

// import "fmt"

// func main() {
// 	text := "Rajesh"
// 	fmt.Printf("%*s\n", 40, "text")
// 	fmt.Printf("%040s\n", "text")
// 	fmt.Printf("%*s\n", 40, text)
// 	fmt.Printf("%040s", text)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	k := '2'
// 	ch := 'A'
// 	str := "B"
// 	fmt.Printf("Type of k : %T\n", k)
// 	fmt.Printf("Type of ch : %T\n", ch)
// 	fmt.Printf("Type of str: %T\n", str)

// }

//ABSOLUTE VALUE OF FLOAT NUMBER

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var val float64 = -19.25
// 	fmt.Printf("Absolute value of %f", val, math.Abs(val)) //the value to be passed Abs function is float

// }

//FINDING LARGEST NUMBER BETWEEN TWO NUMBERS SING math.Max()

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var num1 float64 = 10.23345
// 	var num2 float64 = 20.122
// 	var large float64 = 0
// 	large = math.Max(num1, num2)
// 	fmt.Printf("largest number is : %f", large)
// }

//FINDING SMALLEST NUMBER BETWEEN TWO NUMBERS USING math,Min()

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var num1 float64 = 10.23345
// 	var num2 float64 = 20.122
// 	var small float64 = 0
// 	small = math.Min(num1, num2)
// 	fmt.Printf("Smaller number is : %f", small)
// }

//POWER OF A NUMBER USING THE Pow() FUNCTION

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var num float64 = 10.0000
// 	var p float64 = 3.0
// 	var result float64 = 0
// 	result = math.Pow(num, p)
// 	fmt.Printf("Result: %f", result)
// }

//GOLANG PROGRAM TO DEMONSTRATE THE math.Ceil() function

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var num1 float64 = 10.23345
// 	var num2 float64 = 22.885122
// 	var result float64 = 0

// 	result = math.Ceil(num1)
// 	fmt.Printf("Result is : %f", result)

// 	result = math.Ceil(num2)
// 	fmt.Printf("\nResult is : %f", result)
// }

// GOLANG PROGRAM TO DEMONSTRATE THE math.Floor() FUNCTION

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var num1 float64 = 16.23345
// 	var num2 float64 = 22.885122
// 	var result float64 = 0

// 	result = math.Floor(num1)
// 	fmt.Printf("Result is : %f", result)

// 	result = math.Floor(num2)
// 	fmt.Printf("\nResult is : %f", result)
// }

//PRINTING W/O PACKAGE -

package main

func main() {
	println("Using println istead of fmt.Println")
	print("Using print istead of fmt.Print\n")

	print("Using println istead of fmt.Println")
	print(" Using print istead of fmt.Print")
}
