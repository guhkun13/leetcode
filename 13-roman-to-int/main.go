package main

import "fmt"

func main() {

	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))

}
func romanToInt(s string) int {
	// create the map value
	romanVal := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sumVal := 0
	for i := 0; i < len(s); i++ {
		// roman := string(s[i])
		intValue := romanVal[s[i]]
		sumVal += intValue
		if i < len(s)-1 {
			nextVal := romanVal[s[i+1]]
			// fmt.Println(roman, "=", intValue)
			if nextVal > intValue {
				sumVal -= 2 * intValue
			}

		}
		// fmt.Println("sumVal", sumVal)
	}

	return sumVal

}
