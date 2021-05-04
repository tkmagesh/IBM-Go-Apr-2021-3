package main

import (
	"fmt"
	"math"
)

/*
	Perimeter of Rect = 2 * height * width
	Perimeter of Circle = 2 * pi * radius
*/

type Rectangle struct {
	Height, Width float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type ShapeWithArea interface {
	Area() float32
}

func main() {
	rect := Rectangle{Height: 10, Width: 20}
	//PrintRectArea(rect)
	PrintArea(rect)

	circle := Circle{Radius: 10}
	//PrintCircleArea(circle)
	PrintArea(circle)
}

/* func PrintRectArea(rect Rectangle) {
	fmt.Println(rect.Area())
}

func PrintCircleArea(circle Circle) {
	fmt.Println(circle.Area())
} */

func PrintArea(shapeWithArea ShapeWithArea) {
	fmt.Printf("Area = %f\n", shapeWithArea.Area())
}
