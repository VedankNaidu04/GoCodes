// ****************assinging variable to anonymous function
// package main

// import "fmt"

// func main() {
// 	test := func(x int) {
// 		fmt.Println(x)
// 	}
// 	test(8)
// }

// //****************Method 2 - defining a global variable and adding a test varible gives error
// package main

// import "fmt"
// test := func(x int) {
// 	fmt.Println(x)
// }
// func main() {
// 	test(8)
// }

// // **************** inferred global cannot be declared globally
// package main
// import "fmt"
// b := 10
// func main(){
// 	var a int = 10
// fmt.Println(a,b)
// 	}

// // now tring with a static variable - static global declaration outside function body works
// package main

// import "fmt"

// var a int = 10

// func main() {
// 	b := 10
// 	fmt.Println(a, b)
//}

// // approach 3 to call an anonymous func
// package main

// import "fmt"

// func main() {
// 	test := func(x int) int {
// 		return x * -1
// 	}(8)
// 	fmt.Println(test)
// }

// package main

// import "fmt"

// func calculate(x int, y int) (int, int) {
// 	return x + y, x - y
// }
// func main() {

// 	sum, difference := calculate(10, 20)
// 	fmt.Print("Addition is %d subtraction is %d\n", sum, difference)
// }

// //************GO program to get date month, year in different variables in Goloang
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	YYYY, MM, DD := time.Now().Date()
// 	fmt.Print("Date :", DD)
// 	fmt.Print("Month :", MM)
// 	fmt.Print("year :", YYYY)
// }

// // ************** How to print indivisual elements of date and time in Golang ?

package main

import (
	"fmt"
	"time"
)

func main() {
	currentDateTime := time.Now()

	day := currentDateTime.Day()
	month := currentDateTime.Month()
	year := currentDateTime.Year()

	hour := currentDateTime.Hour()
	min := currentDateTime.Minute()
	sec := currentDateTime.Second()

	fmt.Printf("the report was generated on %d-%02d-%02d at %02d:%02d:%02d", year, month, day, hour, min, sec) // adjusted the format
}
