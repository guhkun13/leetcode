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
	fmt.Println(x)
	toStr := strconv.Itoa(x)
	arrStr := []string{}
	for _, v := range toStr {
		arrStr = append(arrStr, string(v))
	}
	fmt.Println(arrStr)

	for idx := range arrStr {
		if arrStr[idx] != arrStr[len(arrStr)-1-idx] {
			return false
		}
	}

	return true
}
