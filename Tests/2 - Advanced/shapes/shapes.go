package shapes

import (
	"math"
)

type GeometricShape interface {
	getName() string
	area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
	Name   string
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width

}
func (r Rectangle) GetName() string {
	return r.Name

}

type Circle struct {
	Radius float64
	Name   string
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)

}

func (c Circle) GetName() string {
	return c.Name

}
