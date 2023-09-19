package main

import (
	"fmt"
	"strings"
)

func main() {

	input := []string{"f", "flow", "flight"}
	MainFunc(input)
}

func MainFunc(s []string) string {
	fmt.Println(s)
	pref := s[0]
	fmt.Println(pref)

	for i := 1; i < len(s); i++ {
		for !strings.HasPrefix(s[i], pref) {
			pref = pref[:len(pref)-1]
			fmt.Println(pref)
		}
		fmt.Println(s[i])
	}
	fmt.Println(pref)
	return pref
}
