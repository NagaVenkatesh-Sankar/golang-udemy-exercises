package main

import (
	"fmt"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

// receiver of the square type, that implements the shape interface
func (s square) area() float64 {
	return s.length * s.width
}

// receiver of the circle type, that implements the shape interface
func (c circle) area() float64 {
	var a float64
	a = 3.14 * c.radius * c.radius
	return a
}

//interface
type shape interface {
	area() float64
}

// function that accepts any shape
func info(s shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area :", s.area())
}

func main() {
	// create a square
	sq := square{
		length: 4,
		width:  3,
	}
	// print the area
	info(sq)

	// create a circle
	ci := circle{
		radius: 7,
	}
	// print the area
	info(ci)

}
