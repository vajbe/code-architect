package main

import "log"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Reactangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (s Reactangle) Area() float64 {
	return s.Height * s.Width
}

func (r Reactangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func printShapeInfo(s Shape) {
	log.Println("Area: ", s.Area(), " \t\tPerimeter: ", s.Perimeter())
}

func InterExample() {
	r := Reactangle{Width: 10, Height: 10}
	printShapeInfo(r)

	c := Circle{Radius: 23}
	printShapeInfo(c)
}
