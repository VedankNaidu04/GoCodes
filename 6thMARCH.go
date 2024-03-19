//6_MARCH '24
//Copy Golang Slice

package main

import "fmt"

func main() {
	// create two slices
	primeNumbers := []int{2, 3, 5, 7}
	numbers := []int{1, 2, 3}
	// copy elements of primeNumbers to numbers
	copy(numbers, primeNumbers)
	// copy(primeNumbers, numbers) //This gives output as [1 2 3]
	// print numbers
	fmt.Println("Numbers", numbers)
}

// only the first 3 elements of printNumbers are copied in the numbers

// ****************** 2
package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	//
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

// ****************** 3
package main

import "fmt"

func main() {
	//
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//create slice of from index 2 till index 4(5-1)
	intSlice := arr[2:5]

	fmt.Println("Slice elements:")
	for index, ele := range intSlice {
		fmt.Println("index = %d,element = %d\n", index, ele)
	}
	for _, ele := range intSlice {
		fmt.Println(ele)
	}
	for index, _ := range intSlice {
		fmt.Println(index)
	}
}

// ********************* 4
//Sorting a slice of integer in ascending order i Golang
// Golang program to sort slice of integer
// in ascending order.

package main
import "fmt"
import "sort"
func main(){
slice := []int{70, 20, 30, 60, 50, 60, 10, 80, 90, 100}
sort, Ints(slice)
fmt. Printin("Sorted slice:")
for _,ele := range slice {
fmt.Printf("%d."ele)
}
}

// ****************** 5

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"honestly", "is", "the", "best", "Policy"}

	sort.Strings(slice)

	fmt.Println("Sorted Slice :")
	for ele := range slice {
		fmt.Print(ele)
	}
}

// ******************* 6
// Write a function similar to this to check if the slice is sorted or not

// Golang program to check a specified slice
// is sorted or not
package main
import "fmt"
import "sort"
fune main(){ 
var status bool = false
slicel := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} 
slice2 := []int{70, 20, 30, 60, 50, 60, 10, 80, 90, 100}
status = sort.IntsAreSorted(slicel)
if status == true {
fmt. Println("Slicel is sorted")
} else {
fmt. Println("Slicel is not sorted")
}

status = sort.IntsAreSorted(slice2)
if status == true {
fmt.Println("Slice2 is sorted" )
} else {
fmt.Println("Slice2 is not sorted")
}
}
