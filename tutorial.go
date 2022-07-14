package main

import (
	"fmt"
)

func main() {
	var x []int = []int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Print(x[:cap(x)])
	fmt.Println(s[:cap(s)])
}
