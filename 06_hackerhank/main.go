package main

import (
	"fmt"
)

func main() {

	var input string

	input = "saveChangesInTheEditor"

	fmt.Println(input, "=", camelCase(input), "word(s)")

	input = "abcdefxyz"
	fmt.Println(input, "=", caesarCipher(input, 300), "cypher")
}

func camelCase(input string) int32 {

	var words int32
	words = 1
	for _, v := range input {
		if v >= 'A' && v <= 'Z' {
			words++
		}
	}

	return words
}

func caesarCipher(s string, k int32) string {

	var output string

	for _, v := range s {
		char := v
		switch {
		case v >= 'a' && v <= 'z':
			char = char + k
			char = cypherRune(char, 'a', 'z')
		case v >= 'A' && v <= 'Z':
			char = char + k
			char = cypherRune(char, 'A', 'Z')
		}
		output = output + fmt.Sprintf("%c", char)
	}

	return output
}

func cypherRune(c, first, last rune) rune {
	if c > last {
		c--
		c = first + (c - last)
		c = cypherRune(c, first, last)
	}

	return c
}
