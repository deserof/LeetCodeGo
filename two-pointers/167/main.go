package main

import (
	"fmt"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
			continue
		}

		if sum < target {
			left++
			continue
		}

		if sum == target {
			left++
			right++
			return []int{left, right}
		} else {
			break
		}
	}

	return []int{}
}
