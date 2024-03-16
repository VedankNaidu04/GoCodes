// package main

// import "fmt"

// func countDown(number int) {
// 	//
// 	fmt.Println(number)
// 	countDown(number - 1)
// }
// func main() {
// 	fmt.Println("Countdoen Starts:")
// 	countDown(3)
// }

// **************** RECURSIVE FUNCTION WITH CONDITIONAL STATEMENT

package main

import "fmt"

func CountDown(Number int) {

	if Number > 0 {

		fmt.Println(Number)
		//recursive call
		CountDown(Number - 1)
	} else {
		// ends
		fmt.Println("Coundown stops\n")

	}
}
func main() {
	fmt.Println("CounDown starts")
	CountDown(3)
	fmt.Println("CounDown starts")
	CountDown(10)
}
