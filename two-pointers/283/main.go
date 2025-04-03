package main

import (
	"fmt"
)

// time O(n)
// space O(1)

func main() {
	//moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{1, 3, 0, 5})
}

func moveZeroes(nums []int) {
	left := 0

	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	fmt.Println(nums)
}
