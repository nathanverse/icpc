package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)

	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	isMatch := true
	for i := len(s1) - 1; i >= 0; i-- {
		if s1[i] != s3[i+len(s2)] {
			isMatch = false
		}
		dp[i][len(s2)] = isMatch
	}

	isMatch = true
	for i := len(s2) - 1; i >= 0; i-- {
		if s2[i] != s3[i+len(s1)] {
			isMatch = false
		}
		dp[len(s1)][i] = isMatch
	}

	dp[len(s1)][len(s2)] = true
	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}

			if s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}

	return dp[0][0]
}
func main() {
	fmt.Println(isInterleave("", "", ""))
}
