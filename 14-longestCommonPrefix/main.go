package main

func main() {

	input := []string{"flower", "flow", "flight"}
	MainFunc(input)
}

func MainFunc(strs []string) string {
	shortestWord := strs[0]

	// find the shortest words
	for _, v := range strs {
		if len(v) < len(shortestWord) {
			shortestWord = v
		}
	}

	// map all words count
	mp := map[string]int{}
	for _, v := range strs {
		for i := 1; i <= len(shortestWord); i++ {
			mp[v[:i]]++
		}
	}

	lcp := ""
	for k, v := range mp {
		if len(k) > len(lcp) && v >= len(strs) {
			lcp = k
		}
	}

	return lcp

}
