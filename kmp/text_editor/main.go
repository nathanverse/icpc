package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func KMPSearch(lps []int, s string, p string) int {
	i, j := 0, 0
	count := 0
	for i < len(s) {
		if s[i] == p[j] {
			i++
			j++

			if j < len(p) {
				continue
			}
		}

		if j == len(p) {
			count++
		}

		if j <= 0 {
			i++
			continue
		}

		j = lps[j-1]
	}

	return count
}

/*
Binary search approach with simple idea that, we use KMPSearch to search
for certain length, called [mid]. If
+ The occurrences > minAppear, we move [left] upward, record current satisfied result
+ Otherwise, move [right] downward.

=> Complexity: O((n+m)*log(n))
*/
func solve1(text string, pattern string, minAppear int) string {
	left, right := 0, len(pattern)-1
	lps := buildLPS(pattern)
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		count := KMPSearch(lps, text, pattern[0:mid+1])
		if count < minAppear {
			right = mid - 1
		} else {
			res = mid
			left = mid + 1
		}
	}

	if res == -1 {
		return "IMPOSSIBLE"
	}

	return pattern[0 : res+1]
}

/*
*
This approach builds countArr, in which countArr[i] represent the frequency
of prefix pattern_0...[pattern]_i that appears in [text]. CountArr is build
incrementally using kmp approach with two ideas:
+ Everytime j increments, this mean the pattern_0...pattern_j map some string in
text, so => increment countArr[j]
+ As j moves downward due to current mismatch between text[i] and pattern[j], new updated
j represent the new potential length. Because the prefix of length j-1 in pattern is guaranteed
to match with some string in [text]. So we increment countArr[j-1].
  - Because the prefix of length j-1 matches, prefixes with length from 0 to j-1 also match. We
    increment them as well

=> Worst case in O((m+n)*n)
*/
func solve(text string, pattern string, minAppear int) string {
	lps := buildLPS(pattern)
	countArr := make([]int, len(pattern))

	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			countArr[j]++
			j++
			i++

			if j < len(pattern) {
				continue
			}
		}

		if j <= 0 {
			i++
			continue
		}

		j = lps[j-1]
		if j > 0 {
			for k := 0; k < j; k++ {
				countArr[k]++
			}
		}
	}

	for j > 0 {
		j = lps[j-1]
		for k := 0; k < j; k++ {
			countArr[k]++
		}
	}

	for k := len(countArr) - 1; k >= 0; k-- {
		if countArr[k] >= minAppear {
			return pattern[0 : k+1]
		}
	}

	return "IMPOSSIBLE"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 1024, 200000), 200000)
	scanner.Scan()
	text := scanner.Text()
	scanner.Scan()
	pattern := scanner.Text()
	scanner.Scan()
	minLength, _ := strconv.Atoi(scanner.Text())
	fmt.Println(solve1(text, pattern, minLength))
}
