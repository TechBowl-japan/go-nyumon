package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle := Rectangle{
		Width:  30,
		Height: 15.5,
	}

	circle := Circle{
		Radius: 10,
	}

	printArea(rectangle)
	printArea(circle)
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func printArea(s Shape) {
	fmt.Println("面積: ", s.Area(), "㎠")
}

type Shape interface {
	Area() float64
}

func (v Rectangle) Area() float64 {
	return v.Width* v.Height
}

func (v Circle) Area() float64 {
	return v.Radius*v.Radius*math.Pi
}