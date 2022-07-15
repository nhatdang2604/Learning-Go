package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	width  float32
	height float32
}

func (rect Rectangle) area() float32 {
	return rect.width * rect.height
}

func (rect Rectangle) getHeight() float32 {
	return rect.height
}

func (circ Circle) area() float32 {
	return math.Pi * circ.radius * circ.radius
}

func main() {
	circ := Circle{4.5}
	rect := Rectangle{5, 7}

	shapes := []Shape{circ, rect}

	fmt.Println(circ)
	fmt.Println(rect)
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	//Type conversion
	var test Rectangle = shapes[1].(Rectangle)
	fmt.Println(test.getHeight())
}
