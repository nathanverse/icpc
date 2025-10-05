package main

import "fmt"

func isPalindrome(arr string) bool {
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partitionRecur(s string, pos int, state []string, res *[][]string) {
	if pos >= len(s) {
		resState := make([]string, len(state))
		copy(resState, state[:])
		resAdd := append(*res, resState)
		*res = resAdd
		return
	}

	statePos := pos
	for i := pos + 1; i <= len(s); i++ {
		if isPalindrome(s[pos:i]) {
			state = append(state, s[pos:i])
			pos = i
			partitionRecur(s, pos, state, res)
			state = state[:len(state)-1]
			pos = statePos
		}
	}

	return
}

func partition(s string) [][]string {
	res := make([][]string, 0)
	state := make([]string, 0)
	partitionRecur(s, 0, state, &res)

	return res
}

func main() {
	fmt.Println(partition("cbbbcc"))
}
