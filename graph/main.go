package main

import "fmt"

type Node struct {
	Str      string
	Distance int
}

func transformString(s, d string, dict map[string]struct{}) int {
	queue := make([]Node, 0)
	queue = append(queue, Node{Str: s, Distance: 0})
	delete(dict, s)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Str == d {
			return node.Distance
		}

		nodeStr := node.Str
		for i := 0; i < len(s); i++ {
			for char := 'a'; char <= 'z'; char++ {
				if char == rune(nodeStr[i]) {
					continue
				}

				prefixStr := ""
				if i > 0 {
					prefixStr = nodeStr[:i]
				}

				suffixStr := ""
				if i < len(nodeStr)-1 {
					suffixStr = nodeStr[i+1:]
				}

				newStr := prefixStr + string(char) + suffixStr
				if _, ok := dict[newStr]; ok {
					queue = append(queue, Node{Str: newStr, Distance: node.Distance + 1})
					delete(dict, newStr)
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Print(transformString("cat", "dog", map[string]struct{}{
		"bat": {},
		"cot": {},
		"dog": {},
		"dag": {},
		"dot": {},
		"cat": {},
	}))
}
