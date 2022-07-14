package main

import (
	"fmt"
)

func main() {
	var x []int = []int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	b := append(s, 10)
	b[0] = 123
	fmt.Print(x[:cap(x)])
	fmt.Println(s[:cap(s)])
	fmt.Println(b)

	a := make([]int, 5)
	fmt.Println(a)
}
