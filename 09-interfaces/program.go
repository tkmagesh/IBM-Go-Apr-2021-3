package main

import (
	"fmt"
	"interfaces-demo/contracts"
	"interfaces-demo/models"
)

/*
	Perimeter of Rect = 2 * height + 2 * width
	Perimeter of Circle = 2 * pi * radius
*/

func main() {
	rect := models.Rectangle{Height: 10, Width: 20}
	//PrintRectArea(rect)
	/*
		PrintArea(rect)
		PrintPerimeter(rect)
	*/
	fmt.Println(10)
	fmt.Println(rect)
	PrintDimension(rect)

	circle := models.Circle{Radius: 10}
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

func PrintArea(shapeWithArea contracts.ShapeWithArea) {
	fmt.Printf("Area = %f\n", shapeWithArea.Area())
}

func PrintPerimeter(shapeWithPerimeter contracts.ShapeWithPerimeter) {
	fmt.Printf("Perimter = %f\n", shapeWithPerimeter.Perimeter())
}

func PrintDimension(dimension contracts.Dimension) {
	fmt.Printf("Area = %f\n", dimension.Area())
	fmt.Printf("Perimter = %f\n", dimension.Perimeter())
}
