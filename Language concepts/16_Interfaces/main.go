package main

import (
	"fmt"
	"math"
)

type geometricShape interface {
	getName() string
	area() float64
}

type rectangle struct {
	height float64
	width  float64
	name   string
}

func (r rectangle) area() float64 {
	return r.height * r.width

}
func (r rectangle) getName() string {
	return r.name

}

type circle struct {
	radius float64
	name   string
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)

}

func (c circle) getName() string {
	return c.name

}

func printArea(gs geometricShape) {
	fmt.Printf("The area of the %s is %0.2f \n", gs.getName(), gs.area())
}

func main() {

	fmt.Println("Interfaces!")

	r := rectangle{10, 15, "Rectangle"}
	printArea(r)

	c := circle{10, "Circle"}
	printArea(c)

}
