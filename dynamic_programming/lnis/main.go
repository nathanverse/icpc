package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func searchFirstLower(nums []int, l, r, gnum int) int {
	resultIndex := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < gnum {
			resultIndex = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return resultIndex
}

func lengthOfLNIS(nums []int) int {
	dp := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(dp) == 0 || nums[i] <= dp[len(dp)-1] {
			dp = append(dp, nums[i])
			continue
		}

		firstLowerIndex := searchFirstLower(dp, 0, len(dp)-1, nums[i])
		dp[firstLowerIndex] = nums[i]
	}

	return len(dp)
}

func main() {
	const maxCapacity = 64 * 1024 // 64KB
	buf := make([]byte, 0, 2_000_000)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(buf, 2_000_000)
	testCaseNum := 1

	for {
		var heights []int
		var input int
		var err error

		// Read missile heights for one test case
		for {
			scanner.Scan()
			input, err = strconv.Atoi(scanner.Text())
			if err != nil {
				// End of file or invalid input, break the loop
				return
			}
			if input == -1 {
				break
			}
			heights = append(heights, input)
		}

		// Check for the end of all test cases
		if len(heights) == 0 {
			break
		}

		// Find the longest non-increasing subsequence
		longest := lengthOfLNIS(heights)

		// Print the output for the current test case
		fmt.Printf("Test #%d:\n", testCaseNum)
		fmt.Printf("  maximum possible interceptions: %d\n", longest)
		fmt.Println()

		testCaseNum++
	}
}
