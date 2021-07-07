package main

import (
	"fmt"
	"math"
)

type rect struct {
	length  int32
	breadth int32
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return float64(r.length) * float64(r.breadth)
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}
func main() {
	c1 := circle{5.5}

	r1 := rect{5, 6}

	fmt.Println(c1.area(), r1.area())

	shape := []shape{c1, r1}

	for _, shape := range shape {
		fmt.Println(shape.area())

	}
}
