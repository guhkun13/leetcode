package main

import (
	"sort"
)

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	RemoveElement(nums, val)
}

func RemoveElement(nums []int, val int) int {
	count := len(nums)
	for i, v := range nums {
		if v == val {
			nums[i] = 51
			count--
		}
	}

	sort.Ints(nums)
	return count
}
