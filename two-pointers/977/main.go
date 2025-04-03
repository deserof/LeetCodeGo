package main

import (
	"fmt"
	"math"
)

// time O(n)
// space O(n)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	left := 0
	right := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if math.Abs(float64(nums[right])) > math.Abs(float64(nums[left])) {
			res[i] = nums[right] * nums[right]
			right--
		} else {
			res[i] = nums[left] * nums[left]
			left++
		}
	}

	return res
}
