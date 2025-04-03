package main

import (
	"fmt"
	"math"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                // 3
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) // 5
	fmt.Println(longestSubarray([]int{1, 1, 1}))                   // 2
}

func longestSubarray(nums []int) int {
	k := 1
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

	return int(result) - 1
}
