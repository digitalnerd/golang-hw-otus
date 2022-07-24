package main

import (
	// "fmt"
	"fmt"
	"strings"
)
// look at this example -> https://www.geeksforgeeks.org/how-to-remove-duplicate-values-from-slice-in-golang/
func Top10(str string) []string {
	// fmt.Println(str)

	words := strings.Fields(str)
	topWords := make(map[string]bool)
	list := []string{}

	for _, entry := range words {
		if _, value := topWords[entry]; !value {
			topWords[entry] = true
			list = append(list, entry)
		}
	}

	// fmt.Println(reflect.TypeOf(words))

	// for i, word := range words {
	// 	topWords[word] = i
	// }

	// for i, key := range topWords {
	// 	fmt.Println(i, key)
	// }
	// fmt.Println(topWords)

	// for i, word := range words {
	// 	// _, ok := topWords[word]
	// 	if _, ok := topWords[i]; ok {
	// 		// fmt.Println(v)
	// 	}
	// }

	return list
}

func main() {
	str := "one two three one two four five six five"
	fmt.Println(Top10(str))
}
