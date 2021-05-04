package main

import (
	"fmt"
	"math"
)

/*
	Perimeter of Rect = 2 * height + 2 * width
	Perimeter of Circle = 2 * pi * radius
*/

type Rectangle struct {
	Height, Width float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2*r.Height + 2*r.Width
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Shape : RECTANGLE  Height : %f, Width %f", r.Height, r.Width)
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Shape : CIRCLE  Radius : %f", c.Radius)
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

/*
type Dimension interface {
	Area() float32
	Perimeter() float32
}
*/

type Dimension interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func main() {
	rect := Rectangle{Height: 10, Width: 20}
	//PrintRectArea(rect)
	/*
		PrintArea(rect)
		PrintPerimeter(rect)
	*/
	fmt.Println(10)
	fmt.Println(rect)
	PrintDimension(rect)

	circle := Circle{Radius: 10}
	//PrintCircleArea(circle)
	/*
		PrintArea(circle)
		PrintPerimeter(circle)
	*/
	fmt.Println(circle)
	PrintDimension(circle)
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

func PrintPerimeter(shapeWithPerimeter ShapeWithPerimeter) {
	fmt.Printf("Perimter = %f\n", shapeWithPerimeter.Perimeter())
}

func PrintDimension(dimension Dimension) {
	fmt.Printf("Area = %f\n", dimension.Area())
	fmt.Printf("Perimter = %f\n", dimension.Perimeter())
}
