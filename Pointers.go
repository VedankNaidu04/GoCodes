// POINTERS

// Pointer
// • A pointer is a variable that stores the memory address of another variable.
// • memory addresses are in the hexadecimal format, for example, Ox11A3
// • In Go, we can access the memory address using the & operator

// Declaring a pointer
// • There are two important operators that are associated with pointers:
// • * Operator: This is also called a dereferencing operator It is used for the declaration of a pointer, as well as for accessing the value stored at the address that the pointer is pointing to.
// • & Operator: This is known as an address operator. I is used to get the address of a variable.

// The syntax for declaring a pointer in Golang is: var name *Type
// name: This is the name of the pointer variable.
// Type: This is the data type of the variable that the pointer points to.

// // *******************
// package main

// import "fmt"

// //a normal variable whose address the pointer will store
// func main() {
// 	var intData = 20
// 	// declaration of a pointer

// 	var intPointer *int

// 	//intPointer now points towards intData

// 	intPointer = &intData
// 	fmt.Println("what intData stores:", intData)
// 	fmt.Println("address of intData:", &intData)

// 	fmt.Println("what intPointer stores:", intPointer)
// 	//this updates the value of intData using dereferncing operator
// 	*intPointer = 30
// 	fmt.Println("what intData now stores:", intData)
// }

// // *************
// package main

// import "fmt"

// func swap(x int, y int) (int, int) {
// 	return y, x
// }

// func main() {
// 	var a int = 10
// 	var b = 20
// 	a, b = swap(a, b)
// 	fmt.Println(a, b)
// }

// ****************

package main

import "fmt"

func swap(x *int, y *int) {
	t := *x
	*x = *y
	*y = t
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}
