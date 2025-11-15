package main

import "fmt"

func Solve(m, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 || j == n-1 {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = dp[i][j+1] + dp[i+1][j]
		}
	}

	return dp[0][0]
}

func main() {
	fmt.Println(Solve(3, 2))
}
