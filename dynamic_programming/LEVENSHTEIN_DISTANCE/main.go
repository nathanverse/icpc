package main

import (
	"fmt"
)

func Solution(a, b string) int {
	dp := make([][]int, len(a)+1)

	for i := 0; i <= len(a); i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 0; i <= len(b); i++ {
		dp[0][i] = i
	}

	for i := 0; i <= len(a); i++ {
		dp[i][0] = i
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j])
			dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j])
			dp[i+1][j+1] = dp[i+1][j+1] + 1
		}

		fmt.Println(dp[i+1])
	}

	return dp[len(a)][len(b)]
}

func main() {
	fmt.Println(Solution("Carthorse", "Orchestra")) // Su - s , sun - {}, su - {}
}
