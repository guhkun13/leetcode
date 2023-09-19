package main

import "fmt"

func main() {

	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3

	// nums1 := []int{4, 5, 6, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 1, 3}
	// n := 3

	// nums1 := []int{2, 1}
	// m := 2
	// nums2 := []int{}
	// n := 0

	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1

	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// append
	nums1 = append(nums1[:m], nums2...)

	// selection sort
	for i, _ := range nums1 {
		minIndex := i
		for j := i + 1; j < len(nums1); j++ {
			if nums1[minIndex] > nums1[j] {
				minIndex = j
			}
		}
		// swap the element
		nums1[i], nums1[minIndex] = nums1[minIndex], nums1[i]
	}
	fmt.Println(nums1)
}
