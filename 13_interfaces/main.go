package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectange struct {
	width, length float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectange) area() float64 {
	return r.width * r.length
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	cirlce := Circle{radius: 5}
	rectange := Rectange{width: 5, length: 10}

	fmt.Printf("CIRLCE AREA: %f\n", getArea(cirlce))
	fmt.Printf("RECTANGLE AREA: %f\n", getArea(rectange))
}
