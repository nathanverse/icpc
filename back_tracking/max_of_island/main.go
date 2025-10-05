package main

import (
	"fmt"
)

var GoneThroughMap = make(map[int]*BST)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (tree *BST) Add(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Value: value}
		return
	}
	tree.Root = insert(tree.Root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}
	return node
}

func (tree *BST) Search(value int) bool {
	return search(tree.Root, value)
}

func search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	} else if value < node.Value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

func isGoneThrough(x, y int) bool {
	if v, ok := GoneThroughMap[y]; ok {
		if (*v).Search(x) {
			return true
		}
	}
	return false
}

func addToMap(x, y int) {
	if GoneThroughMap[y] == nil {
		GoneThroughMap[y] = &BST{Root: nil}
	}

	(*GoneThroughMap[y]).Add(x)
}

func pickk(grid [][]int, area *int, x, y int) {
	relativePosList := [][]int{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}
	for _, rp := range relativePosList {
		nextX := rp[0] + x
		nextY := rp[1] + y
		if nextX >= 0 && nextY >= 0 && nextX < len(grid[0]) && nextY < len(grid) && !isGoneThrough(nextX, nextY) && grid[nextY][nextX] == 1 {
			addToMap(nextX, nextY)
			*area += 1
			pickk(grid, area, nextX, nextY)
		}
	}
}

func maxAreaOfIsland(grid [][]int) int {
	maxxx := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !isGoneThrough(j, i) && grid[i][j] == 1 {
				addToMap(j, i)
				area := 1
				pickk(grid, &area, j, i)
				maxxx = max(maxxx, area)
			}
		}
	}

	return maxxx
}
func main() {
	grid := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}

	fmt.Println(maxAreaOfIsland(grid))
}
