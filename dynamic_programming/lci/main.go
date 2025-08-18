package main

import "fmt"

func searchFirstGreater(nums []int, l, r, gnum int) int {
	resultIndex := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= gnum {
			resultIndex = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return resultIndex
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(dp) == 0 || nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
			continue
		}

		firstGreaterIndex := searchFirstGreater(dp, 0, len(dp)-1, nums[i])
		if nums[i] == dp[firstGreaterIndex] {
			continue
		}
		dp[firstGreaterIndex] = nums[i]
	}

	return len(dp)
}

func main() {
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	// 3, 10,
}
