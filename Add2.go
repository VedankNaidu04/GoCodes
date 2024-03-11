// n:=10; n2:=20; sum = n1 + n2 ( inferred variable (type 2))
// var n1, n2, int ; n1,n2 = 10, 20; sum:= n1+n2 ( type 3 )

package main

import "fmt"

func main() {
	// declaring variables ( static type variable )
	var n1, n2, sum int
	//var city string
	n1 = 10
	n2 = 20
	sum = n1 + n2
	fmt.Println("sum is", sum)
}
