package main

type Node struct {
	Name   int
	Parity int
}

func IsBipartite(edges [][]int) bool {
	// Build graph
	graph := map[int][]int{}
	nonVisited := map[int]bool{}
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = []int{}
		}

		if graph[edge[1]] == nil {
			graph[edge[1]] = []int{}
		}

		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		nonVisited[edge[0]] = true
		nonVisited[edge[1]] = true
	}

	for source, _ := range nonVisited {
		stack := []*Node{
			{
				Name:   source,
				Parity: 0,
			},
		}
		visited := map[int]*Node{}
		delete(nonVisited, source)
		for len(stack) > 0 {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, childNodeName := range graph[popped.Name] {
				if reachedNode, ok := visited[childNodeName]; ok {
					if reachedNode.Parity == popped.Parity {
						return false
					}

					continue
				}

				childNode := &Node{
					Parity: (popped.Parity + 1) % 2,
					Name:   childNodeName,
				}

				visited[childNodeName] = childNode
				stack = append(stack, childNode)
				delete(nonVisited, childNodeName)
			}
		}
	}

	return true
}

func main() { // 1
}
