package main

import (
	"fmt"
	"sync"
)

var (
	adj     map[int][]int
	tin     map[int]int
	tout    map[int]int
	timer   int
	kingPos int = 1

	once sync.Once
)

func main() {
	var numNodes int
	fmt.Scan(&numNodes)

	adj = make(map[int][]int)
	for i := 0; i < numNodes-1; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	once.Do(initTreeProperties)

	var numQueries int
	fmt.Scan(&numQueries)

	for i := 0; i < numQueries; i++ {
		var queryType, x, y int
		fmt.Scan(&queryType, &x, &y)

		if y == x {
			fmt.Println("YES")
			continue
		}

		if queryType == 0 {
			if isAncestor(x, y) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}

			continue
		}

		if y == kingPos {
			fmt.Println("YES")
			continue
		}

		if isAncestor(y, x) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func initTreeProperties() {
	tin = make(map[int]int)
	tout = make(map[int]int)
	timer = 0

	dfs(kingPos, 0, 0)
}

func dfs(u, p, d int) {
	timer++
	tin[u] = timer

	for _, v := range adj[u] {
		if v == p {
			continue
		}
		dfs(v, u, d+1)
	}
	timer++
	tout[u] = timer
}

func isAncestor(u, v int) bool {
	_, uOk := tin[u]
	_, vOk := tin[v]
	if !uOk || !vOk {
		return false
	}
	return tin[u] <= tin[v] && tout[u] >= tout[v]
}
