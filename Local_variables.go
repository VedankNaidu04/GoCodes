// *************** EXAMPLE OF LOCAL VARIABLE

package main

import "fmt"

func addNos() {
	//local variables
	var sum int
	sum = 2 + 4
}
func main() {
	var sum int = 0
	addNos()
	fmt.Println("Sum is ", sum)
}

// .\Local_variables.go:9:6: sum declared and not used
// .\Local_variables.go:14:25: undefined: sum
//Compilation from top to bottom thus both the error are generateed in the top to bottom format
