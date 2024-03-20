package main

import (
	"fmt"
	"math"
)

/* define an interface */
type Shape interface {
	area() float64
}

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a rectangle */
type Rectangle struct {
	width, height float64
}

/* define */
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/* define */
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Println("Circle area : %f\n", getArea(circle))
	fmt.Println("Rectanglr area : %f\n", getArea(rectangle))
}
