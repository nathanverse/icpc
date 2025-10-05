package main

import "fmt"

// https://leetcode.com/problems/jump-game/description/
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	i := 0
	j := 1
	for i+nums[i] < len(nums)-1 {
		maxI := i
		maxRange := i + nums[i]
		for j <= maxRange {
			if j+nums[j] >= maxI+nums[maxI] {
				maxI = j
			}

			j++
		}

		if nums[maxI] == 0 {
			return false
		}

		i = maxI
		j = maxRange + 1
	}

	return true
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
