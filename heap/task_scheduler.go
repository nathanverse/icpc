package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// Len is the number of elements in the collection.
func (h IntHeap) Len() int {
	return len(h)
}

// Swap swaps the elements at indices i and j.
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	taskCounterMap := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		if _, ok := taskCounterMap[tasks[i]]; !ok {
			taskCounterMap[tasks[i]] = 1
			continue
		}

		taskCounterMap[tasks[i]]++
	}

	freqHeap := make(IntHeap, 0)
	for _, v := range taskCounterMap {
		freqHeap = append(freqHeap, v)
	}

	heap.Init(&freqHeap)
	timee := 0
	for freqHeap.Len() > 0 {
		cycle := n + 1
		incomingHeap := make(IntHeap, 0)
		processedCount := 0
		for cycle > 0 && freqHeap.Len() > 0 {
			cycle--
			processedCount++
			curFreq := heap.Pop(&freqHeap).(int)
			curFreq--

			if curFreq > 0 {
				incomingHeap = append(incomingHeap, curFreq)
			}
		}

		if len(incomingHeap) == 0 {
			timee += processedCount
		} else {
			timee += n + 1
		}

		for _, v := range incomingHeap {
			heap.Push(&freqHeap, v)
		}
	}

	return timee
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	fmt.Println(leastInterval(tasks, 3))
}
