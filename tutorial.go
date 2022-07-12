package main

import "fmt"

func main() {
	fmt.Printf("Floating: %07d \r\n", 123)
	var out string = fmt.Sprintf("Floating: %7d", 123)
	fmt.Printf(out)

}
