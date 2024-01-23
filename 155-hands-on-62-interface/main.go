package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s SQUARE) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func shapeArea(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	s1 := SQUARE{
		length: 16,
		width:  18,
	}

	c1 := circle{
		radius: 12.5,
	}

	shapeArea(s1)
	shapeArea(c1)
}
