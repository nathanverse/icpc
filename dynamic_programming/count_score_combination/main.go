package main

import "fmt"

func Solution(elements []int, n int) int {
	dp := make([][]int, len(elements)+1)

	for i := 0; i < len(elements)+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < len(elements)+1; i++ {
		for j := 0; j < n+1; j++ {
			first := dp[i-1][j]

			var second int
			if j < elements[i-1] {
				second = 0
			} else if j > elements[i-1] {
				second = dp[i][j-elements[i-1]]
			} else {
				second = 1
			}

			dp[i][j] = first + second
		}
	}

	return dp[len(elements)][n]
}

func main() {
	fmt.Println(Solution([]int{2, 3, 7}, 12))
	// 1
	// 2, 2,2
	// 2, 2, 1, 1
	// 2,
}
