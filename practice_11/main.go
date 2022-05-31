package main

import (
	"fmt"
	"math"
)

type Square struct {
	Size float64
}

type Circle struct {
	Radius float64
}

func (s Square) Area() float64 {
	return s.Size * s.Size
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s := Square{10}
	c := Circle{20}
	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)
}
