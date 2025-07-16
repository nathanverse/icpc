package main

import (
	"container/list"
	"fmt"
	avl "github.com/emirpasic/gods/trees/avltree"
)

type Point struct {
	x, y int
}

func isOverlap(x, y, x1, y1 int) bool {

}
func buildSegments(allowedSegmentNum int) (map[int][][]int, error) {
	allowedSegments := make(map[int][][]int)
	for i := 0; i < allowedSegmentNum; i++ {
		r, x, y := 0, 0, 0
		_, err := fmt.Scan(&r, &x, &y)
		if err != nil {
			return nil, err
		}

		if _, ok := allowedSegments[r]; !ok {
			allowedSegments[r] = make([][]int, 0)
		}

		allowedSegments[r] = append(allowedSegments[r], []int{x, y})
	}

	// Sort allowedSegments
	for _, segments := range allowedSegments {

	}

	return allowedSegments
}

func main() {
	start, des := &Point{}, &Point{}
	_, err := fmt.Scan(&(start.x), &(start.y), &(des.x), &(des.y))
	if err != nil {
		return
	}

	var allowedSegmentNum int
	_, err = fmt.Scan(&allowedSegmentNum)
	if err != nil {
		return
	}

}
