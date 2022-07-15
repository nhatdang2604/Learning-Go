package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var f0 int = 0
	var f1 int = 1
	var counter int = -1

	return func() int {
		counter++
		if 0 == counter {
			return f0
		}
		if 1 == counter {
			return f1
		}
		if 1 < counter {
			f0, f1 = f1, f0+f1
			return f1
		}

		return -1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
