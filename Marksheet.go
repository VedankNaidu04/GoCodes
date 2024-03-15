// pass on greater than 40% of 3 subjects

package main

import "fmt"

func main() {
	var RN, Msci, Mmat, Meng, Avg int
	var Name string

	fmt.Println("enter name of student:")
	fmt.Scan(&Name)
	fmt.Println("enter Roll no. of student:")
	fmt.Scan(&RN)

	fmt.Println("Enter marks of Science out of 100 :")
	fmt.Scan(&Msci)
	fmt.Println("Enter marks of Maths out of 100 :")
	fmt.Scan(&Mmat)
	fmt.Println("Enter marks of english out of 100 :")
	fmt.Scan(&Meng)

	Avg = (Mmat + Msci + Meng) / 3
	fmt.Printf("the Average marks is : %d\n", Avg, "%")

	if Avg > 40 {
		fmt.Print("the student passed in the examination")
	} else {
		fmt.Print("the student failed in the examinaton")
	}
}
