package main

import (
	"fmt"
	"math"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
	fmt.Println(findMaxAverage([]int{-1}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
	begin := 0
	window_state := 0
	result := math.Inf(-1)
	for end := range nums {
		window_state += nums[end]

		if end-begin+1 == k {
			result = math.Max(result, float64(window_state))
			window_state -= nums[begin]
			begin++
		}
	}

	return float64(result) / float64(k)
}
