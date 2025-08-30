package main

import "fmt"

// Build tree
func buildTree(arr []int, segmentTree []int, left, right, index int) {
	if left == right {
		segmentTree[index] = arr[left]
		return
	}

	mid := left + (right-left)/2
	buildTree(arr, segmentTree, left, mid, 2*index+1)
	buildTree(arr, segmentTree, mid+1, right, 2*index+2)
	segmentTree[index] = segmentTree[2*index+1] + segmentTree[2*index+2]
}

func findSum(segmentTree []int, left, right, from, to, index int) int {
	if left >= from && right <= to {
		return segmentTree[index]
	}

	if right < from || left > to {
		return 0
	}

	mid := left + (right-left)/2
	return findSum(segmentTree, left, mid, from, to, 2*index+1) + findSum(segmentTree, mid+1, right, from, to, 2*index+2)
}

func updateTree(arr []int, segmentTree []int, from, to, index, uIndex, nValue int) {
	if uIndex > to || uIndex < from {
		return
	}

	if from == to {
		segmentTree[index] = nValue
		arr[uIndex] = nValue
		return
	}

	mid := from + (to-from)/2
	if mid >= uIndex {
		updateTree(arr, segmentTree, from, mid, 2*index+1, uIndex, nValue)
	} else {
		updateTree(arr, segmentTree, mid+1, to, 2*index+2, uIndex, nValue)
	}

	segmentTree[index] = segmentTree[2*index+1] + segmentTree[2*index+2]
}

func updateIncrementTree(lazyTree []int, segmentTree []int, from, to, increment int) {

}

func main() {
	arr := []int{1, 5, 9, 2}
	segmentTree := make([]int, 2*len(arr))

	buildTree(arr, segmentTree, 0, len(arr)-1, 0)

	updateTree(arr, segmentTree, 0, len(arr)-1, 0, 2, 3)

	fmt.Println(arr)
	fmt.Println(segmentTree)
}
