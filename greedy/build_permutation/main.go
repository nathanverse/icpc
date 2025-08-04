package main

import (
	"fmt"
	"sort"
)

func Solve(nums []int) int {
	sort.Ints(nums)

	var cost int
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > i {
			cost += nums[i-1] - i
			continue
		}
		cost += i - nums[i-1]
	}

	return cost
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(Solve(nums))
}
