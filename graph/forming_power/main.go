package main

import (
	"fmt"
	"math/bits"
)

type ExpressionNode struct {
	Presentation            string
	Power                   int
	PossibleElementPowerSet []int
}

func (n *ExpressionNode) NodeKey() string {
	key := fmt.Sprintf("%d:", n.Power)

	for i, v := range n.PossibleElementPowerSet {
		if i != len(n.PossibleElementPowerSet)-1 {
			key += fmt.Sprintf("%d,", v)
			continue
		}

		key += fmt.Sprintf("%d", v)
	}

	return key
}

func shortestPathToExpression(targetPower int) string {
	queue := make([]ExpressionNode, 0)
	queue = append(queue, ExpressionNode{
		Presentation:            "x",
		Power:                   1,
		PossibleElementPowerSet: []int{1},
	})

	visitedNode := map[string]struct{}{}
	for len(queue) > 0 {
		vNode := queue[0]
		queue = queue[1:]

		if vNode.Power == targetPower {
			return vNode.Presentation
		}

		// For the square
		newNodePowerElements := make([]int, len(vNode.PossibleElementPowerSet))
		copy(newNodePowerElements, vNode.PossibleElementPowerSet)
		newNodePowerElements = append(newNodePowerElements, vNode.Power*2)
		newNode := ExpressionNode{
			Presentation:            fmt.Sprintf("(%s)^2", vNode.Presentation),
			Power:                   vNode.Power * 2,
			PossibleElementPowerSet: newNodePowerElements,
		}

		if _, ok := visitedNode[newNode.NodeKey()]; !ok {
			queue = append(queue, newNode)
			visitedNode[newNode.NodeKey()] = struct{}{}
		}

		// For the combination of sub elements
		for _, power := range vNode.PossibleElementPowerSet {
			newNodePowerElements = make([]int, len(vNode.PossibleElementPowerSet))
			copy(newNodePowerElements, vNode.PossibleElementPowerSet)
			newNodePowerElements = append(newNodePowerElements, vNode.Power+power)
			newNode = ExpressionNode{
				Presentation:            fmt.Sprintf("(%s)*(x^%d)", vNode.Presentation, power),
				Power:                   vNode.Power + power,
				PossibleElementPowerSet: newNodePowerElements,
			}

			if _, ok := visitedNode[newNode.NodeKey()]; !ok {
				queue = append(queue, newNode)
				visitedNode[newNode.NodeKey()] = struct{}{}
			}
		}
	}

	return ""
}

func ShortestAdditionChain(n int) []int {
	if n <= 1 {
		return []int{1}
	}

	lb := lowerBound(n)      // at least this many multiplications
	ub := binaryMethodLen(n) // safe upper bound (binary method)

	for limit := lb; limit <= ub; limit++ {
		out := []int(nil)
		if dfsID([]int{1}, n, limit, &out) {
			return out
		}
	}
	return nil
}

// Depth-limited DFS over star-chains: each step is last + chain[i] (including i=last → squaring)
func dfsID(chain []int, target, rem int, out *[]int) bool {
	last := chain[len(chain)-1]

	// goal
	if last == target {
		*out = append([]int{}, chain...)
		return true
	}
	// no steps left
	if rem == 0 {
		return false
	}
	// doubling bound: even if we double rem times, can’t reach target → prune
	if last<<rem < target {
		return false
	}

	// Try larger addends first; include squaring when i==len(chain)-1
	for i := len(chain) - 1; i >= 0; i-- {
		next := last + chain[i]
		if next > target || next <= last {
			continue
		}
		// optional tighter prune: from 'next', with rem-1 steps, can we still reach target?
		if (next << (rem - 1)) < target {
			continue
		}

		// capacity-clamp to avoid aliasing
		nextChain := append(chain[:len(chain):len(chain)], next)
		if dfsID(nextChain, target, rem-1, out) {
			return true
		}
	}
	return false
}

// Lower bound: at least ceil(log2 n) multiplies (each step ≤ doubling).
func lowerBound(n int) int {
	// smallest k with 2^k >= n
	if n <= 1 {
		return 0
	}
	return bits.Len(uint(n - 1))
}

// Upper bound: binary method (square every bit, multiply on 1-bits)
func binaryMethodLen(n int) int {
	if n <= 1 {
		return 0
	}
	// #mult = floor(log2 n) + popcount(n) - 1
	return (bits.Len(uint(n)) - 1) + bits.OnesCount(uint(n)) - 1
}

func main() {
	//fmt.Println(shortestPathToExpression(20))
	for _, n := range []int{15, 31, 64, 73, 127, 191} {
		chain := ShortestAdditionChain(n)
		fmt.Printf("%d → %v (len=%d, mults=%d)\n", n, chain, len(chain), len(chain)-1)
	}
}
