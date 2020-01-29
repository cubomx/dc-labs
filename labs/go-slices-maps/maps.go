package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	sentences := strings.Fields(s)
	for _, val := range sentences {
		words[val]++ 
	}
	return words
}

func main() {
	wc.Test(WordCount)
}
