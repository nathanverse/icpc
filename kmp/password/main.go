package main

import "fmt"

// LPS solves this kind of string: ABCABCDABCABCAE
// Once match j part of [ABCABC  -> miss D] to i part of [ABCABC   -> miss A]
// It is able to recover j part efficiently of [ABCA] to i part [ABCA]
func buildLPS(s string) []int {
	ans := make([]int, len(s))
	n := len(s)
	j := 0
	i := 1
	for i < n {
		if s[i] == s[j] {
			j++
			ans[i] = j
		} else {
			if j > 0 {
				j = ans[j-1]
				continue
			}
		}

		i++
	}

	return ans
}

func solve(s string) string {
	if len(s) < 3 {
		return "Just a legend"
	}

	lps := buildLPS(s)
	possibleLength := lps[len(lps)-1]
	for possibleLength > 0 {
		for i := 1; i < len(lps)-1; i++ {
			if lps[i] == possibleLength {
				return s[0:lps[i]]
			}
		}

		possibleLength = lps[possibleLength-1]
	}

	return "Just a legend"
}

func main() {
	var s string
	fmt.Scan(&s)
	////qwertyqwertyqwerty
	//fmt.Println(buildLPS(s))
	fmt.Println(solve(s))
}
