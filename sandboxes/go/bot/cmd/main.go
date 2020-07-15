package main

import "fmt"

// Write a program that creates two custom struct types called triangle and square
// the square type should be a struct with a field called 'sideLength' of type float 64

//the triangle type should be a struct with a field called heaight of type float 64 and a field of type base of type float 64

//both types should have function called getArea that returns the calculated area of the square or triangle

// Area of a triangle = 0.5 * base * height
// area of square = sideLength * sideLength

// Add a shape interface that defines a function called printArea this function should calculate the area of the given
// shape and print it out to the terminal design the interface so that the printArea function can be called with either
// a triangle or a square

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tri := triangle{height: float64(12), base: float64(10)}
	sq := square{sideLength: float64(22)}

	printArea(tri)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
