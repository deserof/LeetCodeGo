package main

import (
	"fmt"
	"math"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	// 6

	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	// 10
}

func longestOnes(nums []int, k int) int {
	begin := 0
	window_state := 0
	result := 0.0

	for end := range nums {
		if nums[end] == 0 {
			window_state++
		}

		for window_state > k {
			if nums[begin] == 0 {
				window_state--
			}

			begin++
		}

		result = math.Max(result, float64(end-begin+1))
	}

	return int(result)
}
