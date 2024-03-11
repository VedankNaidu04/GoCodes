// out of range values after addition var a int = 255 .....  a = a + 5; will give -251
// package main

// import "fmt"

// func main() {
// 	var a uint8
// 	a = 300
// 	fmt.Printf("a=%d", a)
// }

// where uint8 is unsigned int , unsigned 8-bit integer (0 to 255), better memory management

package main

import "fmt"

func main() {
	var student1 string = "John" //type is static
	var student2 = "jane"        //type is inferred
	x := 2                       // type is infered

	fmt.Println(student1)
	fmt.Printf("student2=%d\n", student2)
	fmt.Println(x)
}
