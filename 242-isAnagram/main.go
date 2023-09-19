package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "mukaa"
	t := "kamuaa"

	// res := isAnagramWithSort(s, t)
	res := isAnagramByCounting(s, t)
	fmt.Println("final res", res)
}

func isAnagramByCounting(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	foundChars := make(map[rune]int, len(s))

	for _, v := range s {
		foundChars[v]++
	}

	for _, v := range t {
		foundChars[v]--
	}

	for _, v := range foundChars {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagramWithSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	ls := strings.Split(s, "")
	lt := strings.Split(t, "")
	// sort
	sort.Strings(ls)
	sort.Strings(lt)

	if strings.Join(ls, "") == strings.Join(lt, "") {
		return true
	}

	return false
}
