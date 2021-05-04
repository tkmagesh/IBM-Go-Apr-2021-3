package contracts

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
