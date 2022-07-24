package main

import (
	"fmt"
	"strings"
)

func Top10(str string) []string {
	fmt.Println(str)

	words := strings.Fields(str)
	for i, w := range words {
		fmt.Println(i, w)
	}
	return nil
}

func main() {
	str := "one two three one two four five six five"
	Top10(str)
}
