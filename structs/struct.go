package main

import "math"

/*
	In Go interface resolution is implicit.
	If the type you pass in matches what the
	interface is asking for, it will compile.
*/

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {

	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
