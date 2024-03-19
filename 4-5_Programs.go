// ****************** 1
package main

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
}

// ****************** 2
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers[:4])
	fmt.Println("Numbers:", numbers[1:4])
	fmt.Println("Numbers:", numbers[1:])
	fmt.Println("Numbers:", numbers[:])
	fmt.Println("Numbers:", numbers[0])

	var num = make([]int, 3, 5)
	num[0] = 10
	printSlice(num)
	println()
	num = append(num, 15)
	printSlice(num)
	println()

}

func printSlice(slice []int) {
	fmt.Printf("len=%d cap=%d slice=%d", len(slice), cap(slice), slice)
}



// ****************** 3
package main

import "fmt"

type Rectangle func(int, int) int

type rectanglePara struct {
	length  int
	breadth int
	color   string
	rect    Rectangle
}

func main() {
	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		rect: func(l, b int) int {
			return l * b
		},
	}
	fmt.Println("Color of the rectangle:", result.color)
	fmt.Println("Area of Rectangle:", result.rect(result.length, result.breadth))

}


// ****************** 4
package main

import "fmt"

type Books struct {
	Title   string
	Author  string
	Subject string
	book_id int
}

func main() {
	var Book1 Books
	// book1 specifications
	Book1.Title = "Go Programming"
	Book1.Author = "Mahesh Kumar"
	Book1.Subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	// book2 specifications
	Book2 := Books{"Telecom Billing", "Zara Ali", "Telecom Billing Tutorial", 6495700}

	// print book1
	printBook(Book1)

	// print book2
	printBook(Book2)

}
func printBook(book Books) {
	fmt.Printf("Book title is %s\n", book.Title)
	fmt.Printf("Book author is %s\n", book.Author)
	fmt.Printf("Book subject is %s\n", book.Subject)
	fmt.Printf("Book book_id is %d\n", book.book_id)
}


// ****************** 5
// struct2 
package main

import "fmt"

type Books struct {
	Title   string
	Author  string
	Subject string
	book_id int
}

func main() {
	var Book1 Books
	// book1 specifications
	Book1.Title = "Go Programming"
	Book1.Author = "Mahesh Kumar"
	Book1.Subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	// book2 specifications
	Book2 := Books{"Telecom Billing", "Zara Ali", "Telecom Billing Tutorial", 6495700}

	// print book1
	fmt.Printf("Book 1 title is %s\n", Book1.Title)
	fmt.Printf("Book 1 author is %s\n", Book1.Author)
	fmt.Printf("Book 1 subject is %s\n", Book1.Subject)
	fmt.Printf("Book 1 book_id is %d\n", Book1.book_id)

	// print book2
	fmt.Printf("Book 2 title is %s\n", Book2.Title)
	fmt.Printf("Book 2 author is %s\n", Book2.Author)
	fmt.Printf("Book 2 subject is %s\n", Book2.Subject)
	fmt.Printf("Book 2 book_id is %d\n", Book2.book_id)

}
