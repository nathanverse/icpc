package main

import "fmt"

type Node struct {
	Name       string
	Distance   int
	ParentNode []*Node
}

func replaceLetter(index int, c rune, word string) string {
	firstPart := ""
	if index != 0 {
		firstPart = word[:index]
	}

	endPart := ""
	if index < len(word)-1 {
		endPart = word[index+1:]
	}

	replacedString := firstPart + string(c) + endPart
	return replacedString
}

func toRes(node *Node, curState []string, index int, res *[][]string) {
	curState[index] = node.Name

	if index == 0 {
		resEle := make([]string, len(curState))
		copy(resEle, curState)
		*res = append(*res, resEle)

		return
	}

	for _, parent := range node.ParentNode {
		toRes(parent, curState, index-1, res)
	}

	return
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	queue := make([]*Node, 0)
	beginNode := &Node{
		Name:       beginWord,
		Distance:   0,
		ParentNode: nil,
	}

	// Storing in map
	dictMap := map[string]struct{}{}
	for _, v := range wordList {
		dictMap[v] = struct{}{}
	}

	queue = append(queue, beginNode)
	visitedNode := map[string]*Node{}
	var resNode *Node
	for len(queue) > 0 {
		popedNode := queue[0]
		queue = queue[1:]

		if popedNode.Name == endWord {
			resNode = popedNode
			break
		}

		for i := 'a'; i <= 'z'; i++ {
			popedName := popedNode.Name
			for j := 0; j < len(popedName); j++ {
				if rune(popedName[j]) == i {
					continue
				}

				childString := replaceLetter(j, i, popedName)

				if _, ok := dictMap[childString]; !ok {
					continue
				}

				if childNode, ok := visitedNode[childString]; !ok {
					childNode = &Node{
						Name:       childString,
						Distance:   popedNode.Distance + 1,
						ParentNode: []*Node{popedNode},
					}
					queue = append(queue, childNode)
					visitedNode[childString] = childNode
				} else if popedNode.Distance < childNode.Distance {
					childNode.ParentNode = append(childNode.ParentNode, popedNode)
				}
			}
		}
	}

	resString := make([][]string, 0)
	if resNode != nil {
		tempConstructedPath := make([]string, resNode.Distance+1)
		toRes(resNode, tempConstructedPath, len(tempConstructedPath)-1, &resString)
	}

	return resString
}

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
