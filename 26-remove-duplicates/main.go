package main

import (
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {

	// use methode like "bubble sort" but only one iteration to do in-place,
	// and move all the duplicates to the last index

	// if len(nums) == 1 {
	// 	return 1
	// }
	lastIndex := 101
	count := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			nums[i-1] = lastIndex
			lastIndex++
			count--
		}
	}

	sort.Ints(nums)
	return count
}
