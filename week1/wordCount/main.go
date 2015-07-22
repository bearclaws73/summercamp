package main

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	counts := map[string]int{}
	for _, word := range strings.Fields(str) {
		counts[word]++

	}
	return counts
}

func main() {
	str := "test test cat"
	result := WordCount(str)
	fmt.Println(result)
}

// x := map[string]int {"time": 4, "test": 6}
// func <name>(<argument>, <argument>...) (<return>, <return>...) {<code>}
//func WordCount(str string) map[string]int
