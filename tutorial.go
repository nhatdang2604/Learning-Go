package main

import "fmt"

type Point struct {
	x int32
	y int32
}

type Circle struct {
	center *Point
	radius float32
}

//Spreading like js
type SpreadingCircle struct {
	*Point
	radius float32
}

func changePoint(point Point) {
	point.x = 123
	point.y = 456
}

func main() {
	var center Point = Point{y: 1, x: 2}
	var circle Circle = Circle{radius: 2, center: &center}
	var spread_circle SpreadingCircle = SpreadingCircle{&Point{1, 2}, 4}
	changePoint(center)
	fmt.Println(circle, center, spread_circle.x, spread_circle.y, spread_circle.radius)
}
