package main

import "fmt"

func Execute(graph map[int][]int) {
	res := make([]int, len(graph))

	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	colors := make([]bool, len(graph))

	for v, _ := range graph {
		for neighbor := range graph[v] {
			if res[neighbor] != -1 {
				colors[res[neighbor]] = true
			}
		}

		for i, color := range colors {
			if !color {
				res[v] = i
				break
			}
		}

		for neighbor := range graph[v] {
			if res[neighbor] != -1 {
				colors[res[neighbor]] = false
			}
		}
	}

	maxNum := -1
	for i := 0; i < len(res); i++ {
		if res[i] > maxNum {
			maxNum = res[i]
		}
	}

	fmt.Println(maxNum + 1)
}

func main() {
	// Read vertexes
	var n, m int
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		panic(err)
	}

	graph := make(map[int][]int)
	for i := 0; i < m; i++ {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			panic(err)
		}
		graph[x-1] = append(graph[x-1], y-1)
		graph[y-1] = append(graph[y-1], x-1)
	}

	Execute(graph)
}
