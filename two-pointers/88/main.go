package main

import (
	"fmt"
)

// time O(n)
// space O(n)

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	//merge([]int{1}, 1, []int{}, 0)
	//merge([]int{0}, 0, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := []int{}

	p1 := 0
	p2 := 0

	for p1 < m && p2 < n {
		if nums1[p1] < nums2[p2] {
			result = append(result, nums1[p1])
			p1++
		} else {
			result = append(result, nums2[p2])
			p2++
		}
	}

	for i := p1; i < m; i++ {
		result = append(result, nums1[i])
	}

	for i := p2; i < n; i++ {
		result = append(result, nums2[i])
	}

	for i := range nums1 {
		nums1[i] = result[i]
	}

	fmt.Println(nums1)
}
