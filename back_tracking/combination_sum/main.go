package main

import (
	"fmt"
	"slices"
)

func compare(arr []int, left, target int) int {
	sum := 0
	for i := left; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum - target
}

func backTrack(arr []int, res *[][]int, curTotal, curPos, pickPos, target int) {
	if len(arr) == 0 || curPos < 0 {
		return
	}

	if curTotal == target {
		tempRes := make([]int, len(arr)-curPos-1)
		copy(tempRes, arr[curPos+1:])

		accessedRes := *res
		accessedRes = append(accessedRes, tempRes)
		*res = accessedRes
		return
	}

	for i := pickPos; i >= 0; i-- {
		if i < len(arr)-1 && arr[i] == arr[i+1] {
			continue
		}

		arr[curPos], arr[i] = arr[i], arr[curPos]

		if curTotal+arr[curPos] <= target {
			backTrack(arr, res, curTotal+arr[curPos], curPos-1, i-1, target)
		}

		arr[curPos], arr[i] = arr[i], arr[curPos]
	}

	return
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	slices.Sort(candidates)
	candidates = append(candidates, 0)
	backTrack(candidates, &res, 0, len(candidates)-1, len(candidates)-2, target)
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 30))
}
