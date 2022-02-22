package helpers

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

// Circle implementing interface Shape's Area method
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}
