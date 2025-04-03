package main

import (
	"fmt"
	"math"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

func minSubArrayLen(target int, nums []int) int {
	begin := 0
	window_state := 0
	result := math.Inf(1)

	for end := range nums {
		window_state += nums[end]

		for window_state >= target {
			window_size := end - begin + 1
			result = math.Min(result, float64(window_size))

			window_state -= nums[begin]
			begin++
		}
	}

	if result == math.Inf(1) {
		return 0
	}

	return int(result)
}
