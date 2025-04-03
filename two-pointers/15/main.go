package main

import (
	"fmt"
	"slices"
)

// time O(n^2)
// space O(1)

func main() {
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{-1,0,1}))
	//fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums) // O(nlogn)

	result := [][]int{}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
