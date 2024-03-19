// *************** Create GO MAP using package make() function
package main

import "fmt"

func main() {
	student := make(map[int]string)

	student[1] = "Harry"
	student[2] = "Lily"
	student[12] = "Hermoine"
	student[7] = "Louise"

	fmt.Println(student)
}

// ****************** Maps of using structure

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.76547, -74.2338},
	"Google":    Vertex{37.8398, -122.0129},
}

func main() {
	fmt.Println(m)

	// Adding a third key-value pair
	m["Nagpur"] = Vertex{67.8778, -77.8876}

	// Printing the map after adding the new pair
	fmt.Println(m)
}

// ****************** Test that a key is present in map with a two value assignment
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// *************** ****** 