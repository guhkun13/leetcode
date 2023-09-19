package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// tc := "A man, a plan, a canal: Panama"
	tc := "0PAP0"

	res := isPalindrome(tc)
	fmt.Println(res)
}

func isPalindrome(s string) bool {

	if len(s) == 0 || len(s) == 1 {
		return true
	}

	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s1 := strings.ToLower(s)
	s2 := strings.Clone(s1)

	// cara 1
	// j := len(s2) - 1
	// for i := 0; i < len(s); i++ {
	// 	if s[i] != s2[j] {
	// 		return false
	// 	}
	// 	j--
	// }

	// cara kedua

	if s1 != reverse(s2) {
		return false
	}

	return true
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
