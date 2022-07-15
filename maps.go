package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	tokens := strings.Split(s, " ")
	var dictionary map[string]int = make(map[string]int, len(tokens))
	for _, token := range tokens {
		dictionary[token]++
	}

	return dictionary
}

func main() {
	wc.Test(WordCount)
}
