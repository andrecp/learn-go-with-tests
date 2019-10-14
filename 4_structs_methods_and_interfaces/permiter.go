package test04

import "math"

// Rectangle is a Struct representing a rectangle
type Rectangle struct {
	width  float64
	height float64
}

// Circle is a struct representing a circle
type Circle struct {
	radius float64
}

// Shape interface is given for structs that implement area
type Shape interface {
	Area() float64
}

// Perimiter calculates the permiter of a rectancle
func (r Rectangle) Perimiter() float64 {
	return r.height*2 + r.width*2

}

// Area calculates the area of a rectancle
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// Area calculates the area of a rectancle
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
