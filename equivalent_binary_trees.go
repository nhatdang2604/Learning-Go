package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if nil == t {
		return
	}

	WalkRecursive(t.Left, ch)
	ch <- t.Value
	WalkRecursive(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var buffer1, buffer2 [10]int

	i := 0
	for v := range ch1 {
		buffer1[i] = v
		i++
	}

	i = 0
	for v := range ch2 {
		buffer2[i] = v
		i++
	}

	for i, v := range buffer1 {
		if v != buffer2[i] {
			return false
		}
	}

	return true
}

func main() {
}
