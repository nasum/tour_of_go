package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	wordArray := strings.Fields(s)
	num := len(wordArray)
	ret := make(map[string]int)
	for i := 0; i < num; i++ {
		(ret[wordArray[i]])++
	}
	return ret
}

func main() {
	wc.Test(wordCount)
}
