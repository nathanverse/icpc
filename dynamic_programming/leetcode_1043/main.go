package main

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr)+1)
	curMax := -1
	for j := len(arr) - 1; j >= 0; j-- {
		if len(arr)-j <= k {
			curMax = max(arr[j], curMax)
			dp[j] = curMax * (len(arr) - j)
			continue
		}

		curMax = arr[j]
		dpJ := arr[j] + dp[j+1]
		for i := 1; i < k; i++ {
			curMax = max(curMax, arr[j+i])
			maxDpI := curMax * (i + 1)
			dpJ = max(maxDpI+dp[j+i+1], dpJ)
		}

		dp[j] = dpJ
	}

	return dp[0]
}
