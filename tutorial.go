package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type something: ")
	// scanner.Scan()
	// input := scanner.Text()
	// fmt.Printf("You typed: %q", input)

	fmt.Printf("Type an integer number: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	fmt.Printf("You typed: %d", input)

}
