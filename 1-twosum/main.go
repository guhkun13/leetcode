package main

import "fmt"

func main() {

	// nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 4}
	target := 6
	// res := twoSumBruteForce(nums, target)
	res := twoSumHashMap(nums, target)
	fmt.Println("final result")
	fmt.Println(res)
}

// brute force
func twoSumBruteForce(nums []int, target int) []int {

	limit := len(nums)
	for ii := 0; ii < limit; ii++ {
		for jj := ii + 1; jj < limit; jj++ {
			if nums[ii]+nums[jj] > target {
				continue
			}
			if nums[ii]+nums[jj] < target {
				continue
			}
			return []int{ii, jj}
		}
	}
	return []int{}

}

func twoSumHashMap(nums []int, target int) []int {

	foundInts := map[int]int{}
	for idx, currentValue := range nums {

		pairValue := target - currentValue

		fmt.Println("existingValue", foundInts)
		fmt.Printf("idx = %d | currentValue = %d, pairValue = %d, target = %d \n \n", idx, currentValue, pairValue, target)

		// check if target - currentValue is exist on hash map
		if _, ok := foundInts[pairValue]; ok {
			foundIdx := foundInts[pairValue]
			fmt.Println("FOUND NEEDED!", idx, foundIdx)
			return []int{idx, foundIdx}
		}
		foundInts[currentValue] = idx
	}

	return []int{}
}
