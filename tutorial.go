package main

import (
	"fmt"
)

func changeFirst(slice []int) {
	slice[0] = 1000
}

func changePrimitiveValue(value int) {
	value = 100
}

func main() {
	var x []int = []int{3, 4, 5}
	y := 1
	changeFirst(x)
	changePrimitiveValue(y)
	fmt.Println(x, y)

}
