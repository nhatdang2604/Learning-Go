package main

import {
	"fmt"
	"math"
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	width float32
	height float32
}

func (rect Rectangle) area() float32 {
	return rect.width * rect.height
}

func (circ Circle) area() float32 {
	return math.Pi * circ.radius * circ.radius;
}


func main() {
	
}
