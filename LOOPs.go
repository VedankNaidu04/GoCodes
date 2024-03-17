// //****************8 for loop

// // package main

// // import "fmt"

// // func main() {
// // 	number := 1
// // 	//loop that runs infinitely
// // 	for {
// // 		//condition to terminate the loop
// // 		fmt.Printf("%d\n", number)
// // 		number++
// // 		if number > 5 {
// // 			break
// // 		}
// // 	}
// // }

// // WHILE LOOP
// // ************* create multiplication table using while loop

// package main

// import "fmt"

// func main() {
// 	// run while loop for 10 times
// 	multiplier := 1
// 	for multiplier <= 10 {
// 		// finding the product
// 		product := 5 * multiplier
// 		// print the multiplication table in format 5 * 1 = 5
// 		fmt.Printf("5 * %d + %d\n", multiplier, product)
// 		multiplier++
// 	}
// }

// // // ********* CONTINUE Statement
// // // the continue statement is used to skip one or more iterations in the loop. it then continues with the next iteration in the loop

// // ************ Range Keyword in  for loop

// package main
// import "fmt"

// func main(){
// 	colors := [] stirng{"Red","Yellow","Green"}
// 	for index, val := range colors {
// 		fmt,Println(index,"-",val)
// 	}
// }

// package main
// import "fmt"
// func main(){
// 	colors := [] stirng{"Red","Yellow","Green"}
// 	for _, val := range colors {
// 		fmt,Println(val)
// 	}
// }

// //NESTED for Loop

// package main

// import "fmt"

// func main() {
//     // Define the dimensions of the grid
//     rows := 3
//     cols := 3

//     // Nested for loop to iterate over rows and columns
//     for i := 0; i < rows; i++ {
//         for j := 0; j < cols; j++ {
//             // Print the coordinates (i, j)
//             fmt.Printf("(%d, %d) ", i, j)
//         }
//         // Print a new line after each row
//         fmt.Println()
//     }
// }

// Program to create an array and prints its elements
package main

import "fmt"

func main() {

	var arrayOfInteger = [5]int{1, 5, 8, 0, 3}
	fmt.Println(arrayOfInteger)
}
// Chaninging array element in go  
package main
import "fmt"
func main(){
	weather := [3]string{"Rainy","Sunny","Cloudy"}

	weather[2]= "stormy"
	fmt.Println(weather)
}

// ****** other ways tp declare a gp array 

package main
import "fmt"
func main(){
// 
	var arrayOfString = [...]string{"Hello","Rajesh"}
	// OR arrayOfString = [...]string{"Hello","Rajesh"}
	fmt.Println(arrayOfString)
}
