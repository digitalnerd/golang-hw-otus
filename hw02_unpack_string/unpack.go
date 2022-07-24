package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {

	var b strings.Builder

	for i, r := range str {
		// the first char is digit
		if i == 0 && unicode.IsDigit(r) == true {
			log.Fatal("String is not correct because of the first char is digit")
		}

		if unicode.IsDigit(r) != true {
			fmt.Fprintf(&b, string(r))
		} else {

			// two digits together
			// if unicode.IsDigit(r+1) == true {
			// 	log.Fatal("String is not correct because of the two digits together")
			// }

			// fmt.Println(str[i+1])

			intVar, err := strconv.Atoi(string(r))
			fmt.Println(intVar)
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Println(&b)

			repl := strings.Repeat(string(str[i-1]), intVar-1)
			fmt.Fprintf(&b, repl)
		}
	}
	fmt.Println(b.String())

	return "", nil
}

func main() {
	// "a4bc2d5e" 	// work
	// "abcd"	 	// work
	// "3abc"	 	// work
	// "a45" 		// work but it check only the first digit
	// "aaa10b"		// doesn't work
	// "aaa1f"		// doesn't work because of the 0 -> intVar-1
	// ""			// work
	Unpack("a4bc2d5e")
}
