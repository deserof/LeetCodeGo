package main

import (
	"fmt"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("b", "c"))
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	p1 := 0
	p2 := 0

	for p1 < len(s) && p2 < len(t) {
		if s[p1] == t[p2] {
			p1++
		}

		p2++
	}

	return p1 == len(s)
}
