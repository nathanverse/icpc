package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return robRecur(root, true, "", &map[string]int{})
}

func getDpIfExist(state string, dp *map[string]int, node *TreeNode) int {
	if v, ok := (*dp)[state]; ok {
		return v
	} else {
		return robRecur(node, true, state, dp)
	}
}

func robRecur(root *TreeNode, getToChoose bool, state string, dp *map[string]int) int {
	if root == nil {
		return 0
	}

	leftState := state + "l"
	rightState := state + "r"

	if getToChoose {
		firstChoice := root.Val + robRecur(root.Left, false, leftState, dp) + robRecur(root.Right, false, rightState, dp)

		leftChoice := 0
		if root.Left != nil {
			leftChoice = getDpIfExist(leftState, dp, root.Left)
		}

		rightChoice := 0
		if root.Right != nil {
			rightChoice = getDpIfExist(rightState, dp, root.Right)
		}

		secondChoice := leftChoice + rightChoice
		(*dp)[state] = max(firstChoice, secondChoice)
		return max(firstChoice, secondChoice)
	} else {
		return getDpIfExist(leftState, dp, root.Left) + getDpIfExist(rightState, dp, root.Right)
	}
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 3}

	// You can now use 'root' to test your solution
	fmt.Println(rob(root))
}
