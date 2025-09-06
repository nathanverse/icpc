package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxx(arr []int) int {
	maxxx := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxxx {
			maxxx = arr[i]
		}
	}

	return maxxx
}

func solve(a, b []int, k int) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	dp := make([][][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([][]int, len(b)+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < len(b)+1; j++ {
			for m := 0; m < k+1; m++ {
				if a[i-1] == b[j-1] {
					dp[i][j][m] = 1 + dp[i-1][j-1][m]
					continue
				}

				a1 := dp[i-1][j][m]
				a2 := dp[i][j-1][m]
				a3 := dp[i-1][j-1][0]
				if m > 0 {
					a3 = 1 + dp[i-1][j-1][m-1]
				}

				dp[i][j][m] = maxx([]int{a1, a2, a3})
			}
		}
	}

	return dp[len(a)][len(b)][k]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)
	for i := range a {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	b := make([]int, m)
	for i := range b {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(solve(a, b, k))
}
