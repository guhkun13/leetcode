package main

import "fmt"

func main() {

	d1 := []int{2, 1, -1}

	fmt.Println(pivotIndex1(d1))
	fmt.Println(pivotIndex2(d1))
}

// This solution, beats 100%. only  need 0ms
func pivotIndex2(nums []int) int {
	leftSum := 0
	rightSum := sumEls(nums[:])

	for idx, v := range nums {
		rightSum -= v
		if leftSum == rightSum {
			return idx
		}
		leftSum += v
		idx++

	}

	return -1
}

func pivotIndex1(nums []int) int {

	for idx := range nums {
		// v = pivot
		leftEl := sumEls(nums[:idx])
		rightEl := sumEls(nums[idx+1:])

		if leftEl == rightEl {
			return idx
		}
	}

	return -1

}

func sumEls(el []int) int {
	res := 0
	for _, v := range el {
		res += v
	}

	return res
}
