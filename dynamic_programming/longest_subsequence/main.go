package main

import "fmt"

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	prevRow := make([]int, len(text2)+1)

	for i := len(text1) - 1; i >= 0; i-- {
		curRow := make([]int, len(text2)+1)
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				curRow[j] = 1 + prevRow[j+1]
			} else {
				curRow[j] = maxx(prevRow[j], curRow[j+1])
			}
		}
		prevRow = curRow
	}

	return prevRow[0]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
