package main

import (
	"fmt"
)

func testBody() int {
	fmt.Println("Inside testBody()")
	return -2
}

func test() int {
	defer fmt.Println("test")
	return testBody()
}

func addAndMinus(x, y int) (add, minus int) {
	add = x + y
	minus = x - y
	return
}

func main() {
	fmt.Println(addAndMinus(12, 13))
	fmt.Println(test())
}
