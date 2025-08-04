package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinimalStrokes(fences []int, l, r, h int) int {
	if r < l {
		return 0
	}

	if r == l {
		return 1
	}

	minHeightIndex := 0
	for i := l; i <= r; i++ {
		if fences[i] < fences[minHeightIndex] {
			minHeightIndex = i
		}
	}

	height := fences[minHeightIndex] - h

	leftStrokes := MinimalStrokes(fences, l, minHeightIndex-1, fences[minHeightIndex])
	rightStrokes := MinimalStrokes(fences, minHeightIndex+1, r, fences[minHeightIndex])

	strokes := height + leftStrokes + rightStrokes
	width := r - l + 1

	return min(strokes, width)
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	fences := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&fences[i])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(MinimalStrokes(fences, 0, n-1, 0))
}
