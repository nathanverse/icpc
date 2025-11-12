package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type WrapperNode struct {
	Node       *TreeNode
	ParentNode *WrapperNode
	Height     int
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	s := make([]*WrapperNode, 0)
	s = append(s, &WrapperNode{
		Node:       root,
		ParentNode: nil,
		Height:     0,
	})

	capturedNodes := make([]*WrapperNode, 2)
	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]

		if node.Node.Val == p.Val {
			capturedNodes[0] = node
		}

		if node.Node.Val == q.Val {
			capturedNodes[1] = node
		}

		if capturedNodes[0] != nil && capturedNodes[1] != nil {
			break
		}

		if node.Node.Left != nil {
			s = append(s, &WrapperNode{
				Node:       node.Node.Left,
				ParentNode: node,
				Height:     node.Height + 1,
			})
		}

		if node.Node.Right != nil {
			s = append(s, &WrapperNode{
				Node:       node.Node.Right,
				ParentNode: node,
				Height:     node.Height + 1,
			})
		}
	}

	if capturedNodes[0] == nil || capturedNodes[1] == nil {
		return nil
	}

	higherHeightNode := capturedNodes[0]
	lowerHeightNode := capturedNodes[1]
	if capturedNodes[1].Height > capturedNodes[0].Height {
		higherHeightNode = capturedNodes[1]
		lowerHeightNode = capturedNodes[0]
	}

	for higherHeightNode.Height != lowerHeightNode.Height {
		higherHeightNode = higherHeightNode.ParentNode
		if lowerHeightNode.Node.Val == higherHeightNode.Node.Val {
			return lowerHeightNode.Node
		}
	}

	for higherHeightNode != nil && lowerHeightNode != nil && higherHeightNode.Node.Val != lowerHeightNode.Node.Val {
		higherHeightNode = higherHeightNode.ParentNode
		lowerHeightNode = lowerHeightNode.ParentNode
	}

	if higherHeightNode == nil {
		return nil
	}

	return higherHeightNode.Node
}
