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

func binary_operator(functor func(int, int) int) func(int, int) int {
	return functor
}

//closure
func test1(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	x := test
	x()
	fmt.Printf("%T\n", x)

	//annonymouse function
	test0 := func() string {
		return "Hello"
	}()

	fmt.Println(test0)
	//test0()

	adder := func(a, b int) int {
		return a + b
	}

	dummy := binary_operator(adder)
	fmt.Println(dummy(2, 3))

	test1("Hê hê")()
}
