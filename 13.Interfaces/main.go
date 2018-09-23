package main

import (
	"fmt"
	"math"
)

func main() {
	c := Cirlce{x: 0, y: 0, radius: 5}
	r := Rectangle{hieght: 10, width: 5}
	fmt.Printf("Circle Area: %f\n", getArea(c))
	fmt.Printf("Rectangle Area: %f\n", getArea(r))

}

// Shape Interface will hold all shapes
type Shape interface {
	area() float64
}

// Cirlce Struct will form a cirlce shape
type Cirlce struct {
	x, y, radius float64
}

// Rectangle struct will form a rectangle shape
type Rectangle struct {
	width, hieght float64
}

//Area method for circle struct
func (c Cirlce) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area method for Rectangle struct
func (r Rectangle) area() float64 {
	return r.hieght * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}
