package main

import (
	"errors"
	"fmt"
	"io"
)

var graph map[int][]int
var isAffectedMap map[int]struct{}
var n, m, d int
var maxDown, maxUp []int

const MAX = 100_000

func readInput() {
	_, err := fmt.Scan(&n, &m, &d)
	if err != nil {
		panic(err)
	}

	isAffectedMap = make(map[int]struct{})
	for i := 0; i < m; i++ {
		var affectedNode int
		_, err := fmt.Scan(&affectedNode)
		if err != nil {
			panic(err)
		}

		if affectedNode < 1 || affectedNode > n {
			panic(err)
		}

		isAffectedMap[affectedNode-1] = struct{}{}
	}

	graph = make(map[int][]int, n) // TODO: check this pre-allocation
	for i := 0; i < n-1; i++ {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		if x < 1 || y < 1 || x > n || y > n {
			panic(err)
		}

		graph[x-1] = append(graph[x-1], y-1)
		graph[y-1] = append(graph[y-1], x-1)
	}
}

func CalMaxDown(src, par int) int {
	if _, ok := isAffectedMap[src]; !ok {
		maxDown[src] = -1
	} else {
		maxDown[src] = 0
	}

	for i := 0; i < len(graph[src]); i++ {
		if graph[src][i] == par {
			continue
		}

		childMaxDown := CalMaxDown(graph[src][i], src)
		if childMaxDown == -1 {
			continue
		}

		if childMaxDown+1 > maxDown[src] {
			maxDown[src] = childMaxDown + 1
		}
	}

	return maxDown[src]
}

func CalMaxUp(src, par int) {
	max1, max2 := -1, -1
	for i := 0; i < len(graph[src]); i++ {
		childNode := graph[src][i]
		if childNode == par {
			continue
		}

		if maxDown[childNode] <= max2 {
			continue
		}

		if maxDown[childNode] <= max1 {
			max2 = maxDown[childNode]
			continue
		}

		max2 = max1
		max1 = maxDown[childNode]
	}

	for i := 0; i < len(graph[src]); i++ {
		childNode := graph[src][i]
		if childNode == par {
			continue
		}
		maxUp[childNode] = max1
		if maxDown[childNode] == max1 {
			maxUp[childNode] = max2
		}

		if maxUp[childNode] > -1 {
			maxUp[childNode] += 2
		}

		if maxUp[src] > -1 && maxUp[src]+1 > maxUp[childNode] {
			maxUp[childNode] = maxUp[src] + 1
		}

		if _, ok := isAffectedMap[childNode]; ok && maxUp[childNode] < 0 {
			maxUp[childNode] = 0
		}

		CalMaxUp(childNode, src)
	}
}

func main() {
	readInput()
	maxDown = make([]int, n)
	maxUp = make([]int, n)
	CalMaxDown(0, -1)
	maxUp[0] = -1
	if _, ok := isAffectedMap[0]; ok {
		maxUp[0] = 0
	}
	CalMaxUp(0, -1)
	res := 0
	for i := 0; i < n; i++ {
		if maxDown[i] <= d && maxUp[i] <= d {
			res += 1
		}
	}

	fmt.Println(res)
}
