// if .... else
package main

import "fmt"

func main() {

	x := 8
	y := 10

	if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater than or equal to y")
	}
}

// EVEN ODD
package main
import "fmt"

func main() {
	var x int
	fmt.Println("enter a number-")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("x is Even")
	} else {
		fmt.Println("x is Odd")
	}
}

// ABSOLUTE VALUE
package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter a number:")
	fmt.Scan(&x)

	if x < 0 {
		fmt.Printf("The absolute value of %d is %d\n", x, -x)
	} else {
		fmt.Printf("The absolute value of %d is %d\n", x, x)
	}
}
