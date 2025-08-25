package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))

}

func isPalindrome(x int) bool {
	toStr := strconv.Itoa(x)
	arrStr := []string{}
	for _, v := range toStr {
		arrStr = append(arrStr, string(v))
	}

	if len(arrStr) == 1 {
		return false
	}

	for idx := range arrStr[:len(arrStr)/2] {
		if arrStr[idx] != arrStr[len(arrStr)-1-idx] {
			return false
		}
	}

	return true
}
