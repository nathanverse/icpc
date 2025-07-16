package main

import (
	"container/heap"
	"fmt"
)

type IntMinHeap []int

// These are the methods that make IntMinHeap satisfy the heap.Interface.
// You write these once for a given element type and ordering.

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] } // For a MIN-heap
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop methods must take a pointer receiver because they modify
// the length of the underlying slice.
func (h *IntMinHeap) Push(x any) {
	// Assert the type of x (it comes as any, but we know it's an int for this heap)
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // Get the last element
	*h = old[0 : n-1] // Truncate the slice
	return x
}

func main() {
	var n, w int
	_, err := fmt.Scan(&n, &w)
	if err != nil {
		return
	}

	graph := make([][]int, n)
	inDegree := make([]int, n)
	// Init graph
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	// Read input
	for i := 0; i < w; i++ {
		var curw int
		_, err := fmt.Scan(&curw)
		if err != nil {
			return
		}

		for j := 0; j < curw; j++ {
			var curSub int
			_, err := fmt.Scan(&curSub)
			if err != nil {
				return
			}

			graph[curSub-1] = append(graph[curSub-1], i)
			inDegree[i]++
		}
	}

	h := &IntMinHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			heap.Push(h, i)
		}
	}

	answer := make([]int, n)
	var lastSub int
	var count int
	for len(*h) > 0 {
		nextNode := heap.Pop(h).(int)
		if count > 0 {
			answer[lastSub] = nextNode + 1
		}
		lastSub = nextNode

		for _, node := range graph[nextNode] {
			inDegree[node]--
			if inDegree[node] == 0 {
				heap.Push(h, node)
			}
		}

		count++
	}

	for _, val := range answer {
		fmt.Println(val)
	}
}
