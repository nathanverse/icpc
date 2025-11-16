package main

import "fmt"

func Solve(n, k int) int {
	dp := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
	}

	for i := n; i >= 0; i-- {
		for j := k; j >= 0; j-- {
			if k-j > n-i {
				dp[i][j] = 0
				continue
			}

			if j == k {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
		}
	}

	return dp[0][0]
}

func main() {
	fmt.Println(Solve(3, 1))
}
