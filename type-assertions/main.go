package main

import (
	"fmt"
	"math"
)

type Celsius float64
type Fahrenheit float64

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func CtoF() {
	var c Celsius = 100.0
	f := Fahrenheit((c * 9 / 5) + 32)

	fmt.Println(c)
	fmt.Println(f)
}

func AreaofShapes() {
	var radius float64
	var length, breadth float64
	fmt.Println("Enter the radius: ")
	fmt.Scan(&radius)
	fmt.Println("Enter the length and breadth: ")
	fmt.Scan(&length, &breadth)
	circle := Circle{Radius: radius}
	rectangle := Rectangle{Width: length, Height: breadth}
	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		fmt.Println("Area: ", shape.Area())
	}
}

func main() {
	AreaofShapes()
}
