package main

// 17FEB
import "fmt"

func main() {
	// const a = 10
	// cannot assign to a (neither addressable nor a map index expression)
	var a int
	a = 10
	a = a + 20
	fmt.Println("a is =", a)
}

//typed constant
// const int  A = 1

//untyped constant
// const A = 1

//Multiple constant declaration
const (
	A int = 1
	B     = 3.14
	C     = "hello"
)
