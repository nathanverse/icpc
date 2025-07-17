package main

import (
	"container/list"
	"fmt"
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	"sort"
)

type Point struct {
	x, y int
}

func isOverlap(x, y, x1, y1 int) bool {
	return x1 <= y
}

func merge(x, y, x1, y1 int) []int {
	res := make([]int, 2)
	res[0] = x

	if y < y1 {
		res[1] = y1
	} else {
		res[1] = y
	}

	return res
}

type Segments [][]int

func (s Segments) Len() int {
	return len(s)
}

func (s Segments) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Segments) Less(i, j int) bool {
	return s[i][0] < s[j][0]
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

	if len(allowedSegments) == 0 {
		return allowedSegments, nil
	}

	// Sort allowedSegments
	for _, segments := range allowedSegments {
		sort.Sort(Segments(segments))
	}

	// Merge segments
	for _, segments := range allowedSegments {
		if len(segments) <= 1 {
			continue
		}

		lastSegment := segments[0]
		mergedSegments := make([][]int, 0)
		mergedSegments = append(mergedSegments, lastSegment)
		for i := 1; i < len(segments); i++ {
			if isOverlap(lastSegment[1], lastSegment[1], segments[i][0], segments[i][1]) {
				mergedSegment := merge(lastSegment[0], lastSegment[1], segments[i][0], segments[i][1])
				lastSegment[0] = mergedSegment[0]
				lastSegment[1] = mergedSegment[1]
				continue
			}

			mergedSegments = append(mergedSegments, segments[i])
			lastSegment = mergedSegments[len(segments)-1]
		}
	}

	return allowedSegments, nil
}

func isAllowed(point Point, allowedSegments map[int][][]int) bool {
	segments, ok := allowedSegments[point.x]
	if !ok {
		return false
	}

}

func main() {
	start, des := &Point{}, &Point{}
	_, err := fmt.Scan(&(start.x), &(start.y), &(des.x), &(des.y))
	if err != nil {
		return
	}

	if start.x == des.x && start.y == des.y {
		fmt.Print(0)
	}

	var allowedSegmentNum int
	_, err = fmt.Scan(&allowedSegmentNum)
	if err != nil {
		return
	}

	allowedSegments, err := buildSegments(allowedSegmentNum)
	if err != nil {
		return
	}

	d := make(map[Point]int)
	q := llq.New()
	q.Enqueue(start)
	for !q.Empty() {
		visitedNode, _ := q.Dequeue()
		visitedPoint, _ := visitedNode.(*Point)
		pushAllowedCells(visitedPoint, allowedSegments)
	}
}
