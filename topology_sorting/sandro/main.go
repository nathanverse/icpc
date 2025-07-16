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
	var n, m int
	_, err := fmt.Scan(&n, &m)
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
	for i := 0; i < m; i++ {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			return
		}

		graph[x-1] = append(graph[x-1], y-1)
		inDegree[y-1]++
	}

	h := &IntMinHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			heap.Push(h, i)
		}
	}

	answer := make([]int, 0)
	for len(*h) > 0 {
		nextNode := heap.Pop(h).(int)
		answer = append(answer, nextNode+1)

		for _, node := range graph[nextNode] {
			inDegree[node]--
			if inDegree[node] == 0 {
				heap.Push(h, node)
			}
		}
	}

	if len(answer) != n {
		fmt.Printf("Sandro fails.")
		return
	}

	for i, val := range answer {
		fmt.Print(val) // Print the number

		// Add a space after each number, except the last one
		if i < len(answer)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println() // Print a newline at the very end
}
