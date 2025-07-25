package main

import (
	"fmt"
	"math"
	"sort"
)

const MAX_DIST = 10001.0 // Use a constant for the "infinity" value

func disc(p1, p2 []int) float64 {
	dx := float64(p1[0] - p2[0])
	dy := float64(p1[1] - p2[1])
	dist := math.Sqrt(dx*dx + dy*dy)

	// This check is part of the problem's specific requirement.
	// Keep it if it's strictly necessary for problem constraints.
	if dist > MAX_DIST {
		return MAX_DIST
	}
	return dist
}

// Solve finds the closest pair of points in the given range [l, r] of the points slice.
// The points slice is assumed to be sorted by X-coordinate.
func Solve(points [][]int, l, r int) float64 {
	numPoints := r - l + 1

	if numPoints <= 1 {
		return MAX_DIST
	}

	if numPoints == 2 {
		return disc(points[l], points[r])
	}

	midIdx := l + (r-l)/2
	midPointX := points[midIdx][0]

	// Recursive calls
	minDistLeft := Solve(points, l, midIdx)
	minDistRight := Solve(points, midIdx+1, r)

	smallerDis := math.Min(minDistLeft, minDistRight)

	tempStrandSet := make([][]int, 0, numPoints) // Pre-allocate with max possible capacity

	// Collect points for the strip.
	// No need to check xDis < smallerDis twice, once is enough.
	for i := l; i <= r; i++ {
		if math.Abs(float64(points[i][0]-midPointX)) < smallerDis {
			tempStrandSet = append(tempStrandSet, points[i])
		}
	}

	// Sort by Y-coordinate
	sort.Slice(tempStrandSet, func(i, j int) bool {
		return tempStrandSet[i][1] < tempStrandSet[j][1]
	})

	for i := 0; i < len(tempStrandSet); i++ {
		for j := i + 1; j < len(tempStrandSet) && (float64(tempStrandSet[j][1]-tempStrandSet[i][1]) < smallerDis); j++ {
			smallerDis = math.Min(smallerDis, disc(tempStrandSet[i], tempStrandSet[j]))
		}
	}

	return smallerDis
}

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			// Handle EOF gracefully if it means end of input
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		if n == 0 {
			break
		}

		points := make([][]int, n)
		for i := 0; i < n; i++ {
			var x, y int
			_, err := fmt.Scan(&x, &y)
			if err != nil {
				panic(err)
			}
			points[i] = []int{x, y}
		}

		// Initial sort by X-coordinate is crucial for the divide-and-conquer strategy.
		sort.Slice(points, func(i, j int) bool {
			return points[i][0] < points[j][0]
		})

		res := Solve(points, 0, len(points)-1)
		if res >= MAX_DIST { // Use >= in case of floating point inaccuracies near MAX_DIST
			fmt.Println("INFINITY")
		} else {
			fmt.Printf("%.4f\n", res)
		}
	}
}
