package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{4, 5, 6}
	var arr2d [][]int = [][]int{{1, 2}, {3, 4}}

	fmt.Println(arr)
	fmt.Println(arr2d)

	sum := []int{0, 0}
	var size int = len(arr2d)
	for i := 0; i < size; i++ {
		sum[0] += arr2d[i][0]
		sum[1] += arr2d[i][1]
	}

	fmt.Println(sum)
}
