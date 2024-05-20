package main

func main() {

}

func IsValid(s string) bool {
	if len(s) <= 0 {
		return false
	}

	if len(s)%2 >= 1 {
		return false
	}

	mapopen := map[string]string{}
	mapopen["("] = ")"
	mapopen["{"] = "}"
	mapopen["["] = "]"

	mapclose := map[string]string{}
	mapclose[")"] = "("
	mapclose["}"] = "{"
	mapclose["]"] = "["

	opener := []string{}

	for _, b := range s {
		val := string(b)
		// if its open, append to openers
		if _, ok := mapopen[val]; ok {
			// fmt.Printf("%s opener \n", val)
			opener = append(opener, val)
		} else {
			if len(opener) == 0 {
				return false
			}

			// its closure. so get the opener value
			valOpen := mapclose[val]
			// lalu cek latest item dari map opener
			latestOpener := opener[len(opener)-1]
			if latestOpener != valOpen {
				return false
			}
			// take out the latest value
			opener = opener[:len(opener)-1]
		}
	}

	return len(opener) == 0
}
