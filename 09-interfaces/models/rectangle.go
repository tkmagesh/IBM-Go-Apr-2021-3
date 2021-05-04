package models

import "fmt"

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
