package main

import (
	"strings"
)

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	tes(s1, s2)
}
func tes(s1 string, s2 string) (res []string) {
	fullstring := s1 + " " + s2
	liststring := strings.Split(fullstring, " ")

	mapvalue := map[string]int{}
	for _, v := range liststring {
		mapvalue[v]++
	}

	uncommons := []string{}
	for word, wordCount := range mapvalue {
		if wordCount == 1 {
			uncommons = append(uncommons, word)
		}
	}
	return uncommons
}
