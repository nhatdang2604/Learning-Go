package main

import (
	"fmt"
)

func changeValue0(value *int) {
	*value = 10
}

func changeString0(value string) {
	value = "hehehehhe"
}

func changeString1(value *string) {
	(*value) += "hoho"
}

func main() {
	x := 7
	str := "hello world"
	changeValue0(&x)
	changeString0(str)
	fmt.Println(x, str, &str)
	changeString1(&str)
	fmt.Println(str, &str)
}
