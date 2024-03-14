//DECISION MAKING / if STATEMENT

package main

import "fmt"

func main() {

	var a, b int = 10, 40
	if a < b {
		fmt.Println("b is greater than a")
	}
	if 18 > 10 {
		fmt.Println("18 is greater than 10")
	}
	if true {
		fmt.Println("staya hai")
	}
	if 20 != 18 {
		fmt.Println("hanji haanji dedo mic pe light mai bolna bohot kuch")
	}
}

IF with statement /declaration

package main

import "fmt"

func main() {
	x := 12
	if y := 10; x > y {
		fmt.Println("x is greater than y")
	}
	y := 20
	fmt.Println(y)
}

// ******** CODE ********
package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	fmt.Println("Enter your name & age - ")

	// "_" is used when we want to use variable but do not want to return the value, if some variable , say a is used in place of _ then it gives warning as a declared and not used
	if _, err := fmt.Scan(&name, &age); err != nil { //on input - vedank naidu; expected integer; exit status 1
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Your name is : \n", name)
	fmt.Println("Your name is : \n", age)
}

//PRINT THE BIGGEST NO. OF THE THREE

package main

import "fmt"

func main() {

	var a, b, c int = 12, 34, 48

	if a < c && b < c {
		fmt.Println("c is largest")
	}
	if a < b && c < b {
		fmt.Println("b is largest")
	}
	if c < a && b < a {
		fmt.Println("a is largest") // nested if - Complexity = O(1)
	}
}


package main
import "fmt"
func main(){

	fmt.Println( "enter 3 numbers")
	var first int
	fmt.Scanln(&first)
	var second int 
	fmt.Scanln(&second)
	var third int 
	fmt.Scanln(&third)

}