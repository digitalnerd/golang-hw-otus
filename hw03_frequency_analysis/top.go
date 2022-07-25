package main

import (
	"fmt"
	"sort"
	"strings"
)

// this one was very helpful -> https://www.geeksforgeeks.org/how-to-sort-golang-map-by-keys-or-values/
func Top10(str string) {

	words := strings.Fields(str)
	topWords := make(map[string]int)

	for _, v := range words {
		if _, ok := topWords[v]; !ok {
			topWords[v] += 1
		} else {
			topWords[v] += 1
		}
	}

	fmt.Println(topWords)

	keys := make([]string, 0, len(topWords))
	for key := range topWords {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return topWords[keys[i]] > topWords[keys[j]]
	})

	for _, k := range keys{
        fmt.Println(k, topWords[k])
    }
}

func main() {
	str := "one one one one one one one one two two three four five five five six six six six six"
	Top10(str)
}
