package main

import (
	"fmt"
	"unicode"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(isPalindrome("ab@a"))
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		leftChar := unicode.ToLower(rune(s[left]))
		rightChar := unicode.ToLower(rune(s[right]))

		if !unicode.IsLetter(leftChar) && !unicode.IsDigit(leftChar) {
			left++
			continue
		}

		if !unicode.IsLetter(rightChar) && !unicode.IsDigit(rightChar) {
			right--
			continue
		}

		if leftChar != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}
