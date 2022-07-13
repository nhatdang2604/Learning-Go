package main

import (
	"fmt"
)

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("You can ride alone")
	} else if age >= 14 {
		fmt.Println("you can ride with a parent")
	} else {
		fmt.Println("You can't ride")
	}

}
