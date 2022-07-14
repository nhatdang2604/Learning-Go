package main

import (
	"fmt"
)

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	delete(mp, "apple")
	mp["pear"] = 1234
	fmt.Println(mp)
	fmt.Println(mp["apple"])

	val, err := mp["apocalypse"]
	fmt.Println(val, err)

	for key, value := range mp {
		fmt.Println(key, value)
	}
}
