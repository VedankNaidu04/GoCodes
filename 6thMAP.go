//map
//• In Go, the map data structure stores elements in key/value pairs. Here, keys are unique identifiers that are associated with each value on a map.
// Program to create a map and print its keys and values
package main

import "fmt"

func main() {
	// create a map
	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "python": 81}
	fmt.Println(subjectMarks)
}

// ACCESS VALUE OF MAP IN GOLANG

package main

import "fmt"

func main() {
	//create map
	flowerColor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Rose": "Red"}

	fmt.Println(flowerColor["Sunflower"])
	fmt.Println(flowerColor["Rose"]) // use the key in the bracket in the [] , not the key
	fmt.Println(flowerColor["White"]) // - this prints blank
}

// ****** Change the Value of a map in Golang

package main
import"fmt"
func main(){

	//create map
	capital := map[string]string{"Nepal": "Kathmandu","US": "NewYork"}
	fmt.Println("Initial Map  : ",capital)

	capital["US"] = "Washington Dc"

	fmt.Println("Updated Map  : ",capital)
}

// **************** Key wise
package main

import "fmt"

func main() {
	// create a map
	students := map[int]string{1: "John", 2: "Lily"}
	fmt.Println("Initial Map: ", students)
	// add element with key 3
	students[3] = "Robin"
	// add element with key 5
	students[5] = "Julie"
	fmt.Println("Updated Map: ", students)
}

// Delete element of Go Map Element

package main

import "fmt"

func main() {

	//create map
	personAge := map[string]int{"Hermoine": 21, "HarryP": 20, "Ron": 25}

	delete(personAge, "Ron")
	fmt.Println("updated Map: ", personAge)
}

// Looping through a map in Golang

package main

import "fmt"

func main() {
	//create a map
	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	for number, squared := range squaredNumber {
		fmt.Printf("Sqaure of %d is %d\n", number, squared)
	}
} // order may or may not change while printing the map keys and values
