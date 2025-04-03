package main

import (
	"fmt"
)

// time O(n)
// space O(n) rune array

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}

func reverseVowels(s string) string {
	runes := []rune(s)

	p1 := 0
	p2 := len(runes) - 1

	for p1 < p2 {
		if runes[p1] != 'A' &&
			runes[p1] != 'a' &&
			runes[p1] != 'E' &&
			runes[p1] != 'e' &&
			runes[p1] != 'I' &&
			runes[p1] != 'i' &&
			runes[p1] != 'O' &&
			runes[p1] != 'o' &&
			runes[p1] != 'U' &&
			runes[p1] != 'u' {
			p1++
			continue
		}

		if runes[p2] != 'A' &&
			runes[p2] != 'a' &&
			runes[p2] != 'E' &&
			runes[p2] != 'e' &&
			runes[p2] != 'I' &&
			runes[p2] != 'i' &&
			runes[p2] != 'O' &&
			runes[p2] != 'o' &&
			runes[p2] != 'U' &&
			runes[p2] != 'u' {
			p2--
			continue
		}

		runes[p1], runes[p2] = runes[p2], runes[p1]
		p1++
		p2--
	}

	return string(runes)
}
