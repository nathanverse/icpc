package main

import "fmt"

func findExistAdjancentRow(matrix [][]int, i, j, target int) [][]int {
	res := make([][]int, 0)
	for _, v := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		if i+v[0] >= len(matrix) || j+v[1] >= len(matrix[0]) || i+v[0] < 0 || j+v[1] < 0 {
			continue
		}

		if target == matrix[i+v[0]][j+v[1]] {
			res = append(res, []int{i + v[0], j + v[1]})
		}
	}

	return res
}

func Solve(matrix [][]int, pattern []int) bool {
	if len(pattern) <= 0 {
		return true
	}

	if len(matrix) <= 0 {
		return false
	}

	if len(matrix[0]) <= 0 {
		return false
	}

	cache := map[string][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == pattern[0] {
				cache[fmt.Sprintf("%d:%d", i, j)] = []int{i, j}
			}
		}
	}

	if len(cache) == 0 {
		return false
	}

	for i := 1; i < len(pattern); i++ {
		newCache := map[string][]int{}
		for _, v := range cache {
			matches := findExistAdjancentRow(matrix, v[0], v[1], pattern[i])
			for _, match := range matches {
				newCache[fmt.Sprintf("%d:%d", match[0], match[1])] = []int{match[0], match[1]}
			}
		}

		cache = newCache
	}

	return len(cache) > 0
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 3, 6},
		{7, 8, 9},
	}

	pattern := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 6, 9}

	fmt.Println(Solve(matrix, pattern))
}
