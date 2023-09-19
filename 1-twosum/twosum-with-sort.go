package main

import "fmt"

func twoSumWithSort(nums []int, target int) []int {
	// sort the nums
	fmt.Println(nums)
	nums1 := sort(nums)
	fmt.Println("sort finished")
	fmt.Println(nums1)

	// split array based on its value
	nums2 := split(nums1, target)
	fmt.Println("split finished")
	fmt.Println(nums2)

	res := findTwo(nums2, target)
	fmt.Println("findTwo finished")
	fmt.Println(res)

	return res
}

func sort(nums []int) []int {
	fmt.Println("sort started")
	i := 0
	for i < len(nums)-1 {
		fmt.Println("i = ", i)
		cur := nums[i]
		next := nums[i+1]
		if cur > next {
			// swap
			nums[i] = next
			nums[i+1] = cur
		}
		i++
	}
	return nums
}

func split(nums []int, target int) []int {
	fmt.Println("split started")
	for idx, i := range nums {
		if i > target {
			return nums[0:idx]
		}
	}
	return nums
}
