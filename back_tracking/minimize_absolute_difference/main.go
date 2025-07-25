package main

import (
	"fmt"
)

func exist(nums []int, l int, i int) bool {
	for j := l; j < i; j++ {
		if nums[j] == nums[i] {
			return true
		}
	}

	return false
}

func solve(nums []int, l int, comp []int) ([]int, []int) {
	if l == 4 {
		upper := nums[0]*nums[3] - nums[1]*nums[2]
		if upper < 0 {
			upper = -upper
		}

		return []int{upper, nums[1] * nums[3]}, comp
	}

	curUpper := -1
	curLower := -1
	var minComp []int
	for i := l; i < len(nums); i++ {
		if exist(nums, l, i) {
			continue
		}

		nums[l], nums[i] = nums[i], nums[l]
		res, nodeComp := solve(nums, l+1, append(comp, i))
		if curUpper == -1 || res[0]*curLower < curUpper*res[1] {
			curUpper = res[0]
			curLower = res[1]
			minComp = nodeComp
		}
		nums[l], nums[i] = nums[i], nums[l]
	}

	return []int{curUpper, curLower}, minComp
}

func main() {
	nums := make([]int, 5)
	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&nums[i])
		if err != nil {
			panic(err)
		}
	}

	comp := make([]int, 0)
	_, minComp := solve(nums, 0, comp)
	for i := 0; i < len(minComp); i++ {
		fmt.Print(minComp[i])
		if i != len(minComp)-1 {
			fmt.Print(" ")
		}
	}
}
