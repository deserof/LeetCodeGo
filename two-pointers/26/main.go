package main

import (
	"fmt"
)

// time O(n)
// space O(1)

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	// output: [0,1,2,3,4,_,_,_,_,_]
}

func removeDuplicates(nums []int) int {

	left := 0
	for right := range nums {
		if nums[left] == nums[right] {
			right++
		} else if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
			right++
		}
	}

	return left + 1
}
