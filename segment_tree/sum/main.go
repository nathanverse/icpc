package main

import "fmt"

type RangeSum struct {
	lazySegTree []int
	segTree     []int
	arr         []int
}

func newRangeSum(arr []int) *RangeSum {
	segmentTree := make([]int, 2*len(arr))
	lazySegTree := make([]int, 2*len(arr))
	buildTree(arr, segmentTree, 0, len(arr)-1, 0)

	return &RangeSum{lazySegTree, segmentTree, arr}
}

func (rs *RangeSum) FindSum(from, to int) int {
	return rs.findSum(0, len(rs.arr)-1, from, to, 0)
}

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

func (rs *RangeSum) findSum(left, right, from, to, index int) int {
	if right < from || left > to {
		return 0
	}

	if rs.lazySegTree[index] != 0 {
		rs.segTree[index] += (right - left + 1) * rs.lazySegTree[index]
		if 2*index+1 < len(rs.lazySegTree) {
			rs.lazySegTree[2*index+1] = rs.lazySegTree[index]
		}

		if 2*index+2 < len(rs.lazySegTree) {
			rs.lazySegTree[2*index+2] = rs.lazySegTree[index]
		}

		rs.lazySegTree[index] = 0
	}

	if left >= from && right <= to {
		return rs.segTree[index]
	}

	mid := left + (right-left)/2
	return rs.findSum(left, mid, from, to, 2*index+1) + rs.findSum(mid+1, right, from, to, 2*index+2)
}

func (rs *RangeSum) UpdateTree(uIndex, nValue int) {
	rs.updateTree(0, len(rs.arr)-1, 0, uIndex, nValue)
}

func (rs *RangeSum) updateTree(from, to, index, uIndex, nValue int) {
	if uIndex > to || uIndex < from {
		return
	}

	if from == to {
		rs.segTree[index] = nValue
		rs.arr[uIndex] = nValue
		return
	}

	mid := from + (to-from)/2
	if mid >= uIndex {
		rs.updateTree(from, mid, 2*index+1, uIndex, nValue)
	} else {
		rs.updateTree(mid+1, to, 2*index+2, uIndex, nValue)
	}

	rs.segTree[index] = rs.segTree[2*index+1] + rs.segTree[2*index+2]
}

func (rs *RangeSum) UpdateIncrementTree(from, to, increment int) {
	rs.updateIncrementTree(0, len(rs.arr)-1, 0, from, to, increment)
}

func (rs *RangeSum) updateIncrementTree(left, right, index, from, to, increment int) {
	if right < from || left > to {
		return
	}

	if rs.lazySegTree[index] != 0 {
		rs.segTree[index] += (right - left + 1) * rs.lazySegTree[index]
		if 2*index+1 < len(rs.lazySegTree) {
			rs.lazySegTree[2*index+1] = rs.lazySegTree[index]
		}

		if 2*index+2 < len(rs.lazySegTree) {
			rs.lazySegTree[2*index+2] = rs.lazySegTree[index]
		}

		rs.lazySegTree[index] = 0
	}

	if left >= from && right <= to {
		rs.segTree[index] += (right - left + 1) * increment
		if 2*index+1 < len(rs.lazySegTree) {
			rs.lazySegTree[2*index+1] = increment
		}

		if 2*index+2 < len(rs.lazySegTree) {
			rs.lazySegTree[2*index+2] = increment
		}

		return
	}

	mid := left + (right-left)/2
	rs.updateIncrementTree(left, mid, 2*index+1, from, to, increment)
	rs.updateIncrementTree(mid+1, right, 2*index+2, from, to, increment)

	rs.segTree[index] = rs.segTree[2*index+1] + rs.segTree[2*index+2]
}

func (rs *RangeSum) Print() {
	fmt.Printf("Lazy segment tree: %v\n", rs.lazySegTree)
	fmt.Printf("Segment tree: %v\n", rs.segTree)
	fmt.Printf("Array: %v\n", rs.arr)
}

func main() {
	arr := []int{3, 5, 8, 10, 4, 9, 3}
	rangeSum := newRangeSum(arr)
	rangeSum.UpdateIncrementTree(0, 3, 3)
	fmt.Printf("Sum: %d\n", rangeSum.FindSum(0, 2))
	rangeSum.Print()
	//fmt.Println(rangeSum.FindSum(1, 3))
}
