package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	separated := strings.Fields(s)
	for _, val := range separated {
		words[val]++ 
	}
	return words
}

func main() {
	wc.Test(WordCount)
}
