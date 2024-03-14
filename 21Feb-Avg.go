package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var avg float64

	num1 = 10
	num2 = 5
	avg = (float64(num1) + float64(num2)) / 2 //cannot use ((num1) + (num2)) / 2 (value of type int) as float64 value in assignment
	fmt.Printf("Average of %d and %d is %.2f\n", num1, num2, avg)
}
