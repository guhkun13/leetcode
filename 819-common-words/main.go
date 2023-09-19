package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// p := "Bob hit a ball, the hit BALL flew far after it was hit."
	// banned := []string{"hit"}

	// p := "a."
	// banned := []string{""}

	// p := "a, a, a, a, b,b,b,c, c"
	// banned := []string{"a"}

	// p := "Bob. hIt, baLl"
	// banned := []string{"bob", "hit"}

	// p := "Z? Z, Z. S. O. z; X, R. k? X, R' M! D! i. W, p. X, t; s, U; T? Z? W! X, O. g, M; y? t; X; O, X' C' Y; x! q! Y' T; u? R. j? w, M. F' n' F; y, V' z. R, V; x' y? F' m' p? M. w, n! Y' Y? i. S' P? w; w; y! Z; P' o? I, H! L; U; p' i; s' Z. V; S! V! H! y' I? K; d. L! r? u! U. O! s? j. y. G, g, r; Z, X, x' L! l? Z, w! Z' W! b. N! T. P! Y, Z. u! Z, q! Y? P' G' u' t, N' k, H' T, I. S' q? J. q! i? E! Q. O, j' r; r' L' C, z! G, p. S. p' s' L! u. S. t, V; z, Z' p! t? Z. x! h; T; T' V, w; P? Q' T! q. J; E? n. V' X. M? Q, v; U; O, H? h; T. s, n! Y? a, N; o, V, o. S! K' j! Y, W! v; Q! u? U' l. k. r, o. o; m. E, I. n! H' w? u? x! w! U' m; w; R' n. y. Y, W' d; P? z! K! g? m, J' i. t, j. x! F. F' U? K! r' V, r, s! O, Q, v, v, c. E. s! X. k; F' Y! r? P! g! r! V! w; S! X, S! N, z? m. y. B; Y' n' U? r. u; R. l? U? v, r, y' W' W' Q; n' Y? Z, L? O. T? Q' q. l, z! V. T. k; x' q! s; u? W! x' X; P; m! t; T? X, v, v' t. Y' q; X? u; V; X! q! w. j! W; z; v. u, j. w; v. z; W' P' z; l! l. o! Z, Y. H; Q; k' O' m, U!"
	// banned := []string{"c", "y", "u", "i", "p", "j", "g", "r", "o", "h", "t", "m", "k", "z", "s", "q", "v", "a", "e", "x"}

	p := "Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. "
	banned := []string{"hit"}

	res := mostCommonWord(p, banned)

	fmt.Println("FINAL RESULT", res)
}
func mostCommonWord(paragraph string, banned []string) string {
	// lower case the paragraph
	// fmt.Println(paragraph)

	// lower
	cleanStr := strings.ToLower(paragraph)
	// fmt.Println(cleanStr)

	// remove banned words if exist
	// fmt.Println("banned", banned, len(banned))
	if len(banned) > 0 && banned[0] != "" {
		strBanned := ""
		idx := 0
		for _, v := range banned {
			if idx > 0 {
				strBanned += `|`
			}
			strBanned += fmt.Sprintf(`\b%s\b`, v)
			idx++
		}
		// fmt.Println("strBanned=", strBanned)

		reBanned := regexp.MustCompile(strBanned)
		cleanStr = reBanned.ReplaceAllString(cleanStr, "")
		fmt.Println("cleanStr=", cleanStr)
	}

	// remove all chars aneh
	// reSymbol := regexp.MustCompile(`[!?',;.]`)
	// cleanStr = reSymbol.ReplaceAllString(cleanStr, " ")
	// fmt.Println(cleanStr)

	// remove white spaces more than 1
	reWhiteSpace := regexp.MustCompile(`[!?',;.|\s]{2,}`)
	cleanStr = reWhiteSpace.ReplaceAllString(cleanStr, " ")
	fmt.Println(cleanStr)

	// remove all trailing space
	cleanStr = strings.TrimSpace(cleanStr)
	fmt.Println(cleanStr)

	mapWords := mapWordCount(cleanStr)
	fmt.Println(mapWords)

	return getMaxCountWord(mapWords)
}

// convert string to list of string
func mapWordCount(in string) map[string]int {
	delimiter := " "
	stringList := strings.Split(in, delimiter)

	result := make(map[string]int, 0)
	for _, word := range stringList {
		result[word]++
	}
	return result
}

func getMaxCountWord(in map[string]int) string {
	res := ""
	maxCount := 1
	idx := 0
	for k, v := range in {
		if idx == 0 {
			res = k
			maxCount = v
		}
		if v > maxCount {
			res = k
			maxCount = v
		}

		idx++
	}
	return res
}
