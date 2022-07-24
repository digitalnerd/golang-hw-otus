package main

import (
	"fmt"
	// "reflect"
	"strings"
)

func Top10(str string) []string {
	fmt.Println(str)

	words := strings.Fields(str)
	topWords := make(map[int]string)
	// fmt.Println(reflect.TypeOf(words))
	for i, word := range words {
		// fmt.Println(word)
		// put words to a map
		// topWords = make(map[int]string)
		topWords[i] = word
		for i, key := range topWords {
			fmt.Println(i, key)
		}

	}
	return nil
}

func main() {
	str := "one two three one two four five six five"
	Top10(str)
}
