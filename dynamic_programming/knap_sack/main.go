package main

import "fmt"

type Stuff struct {
	Weight int
	Value  int
}

func maxx(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func solve(manyStuff []Stuff, capacity int) ([]Stuff, int) {
	if len(manyStuff) < 1 || capacity < 1 {
		panic("Invalid params")
	}

	dp := make([][]int, len(manyStuff)+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	var pickValue, notPickedValue int
	for i := 1; i <= len(manyStuff); i++ {
		for w, _ := range dp[i] {
			curStuff := manyStuff[i-1]
			pickValue = -1
			if w >= curStuff.Weight {
				pickValue = curStuff.Value + dp[i-1][w-curStuff.Weight]
			}

			notPickedValue = dp[i-1][w]
			dp[i][w] = maxx(pickValue, notPickedValue)
		}
	}

	optimizedKnapSack := dp[len(manyStuff)][capacity]
	optimizedStuff := make([]Stuff, 0)
	i, j := len(manyStuff), capacity
	for i > 0 {
		pickItem := manyStuff[i-1]
		if dp[i][j] == pickItem.Value+dp[i-1][j-pickItem.Weight] {
			optimizedStuff = append(optimizedStuff, pickItem)
			i, j = i-1, j-pickItem.Weight
		} else {
			i, j = i-1, j
		}
	}

	return optimizedStuff, optimizedKnapSack
}

func main() {
	stuff := []Stuff{{
		Weight: 1,
		Value:  1,
	}, {
		Weight: 1,
		Value:  2,
	}, {
		Weight: 2,
		Value:  2,
	}, {
		Weight: 6,
		Value:  4,
	}, {
		Weight: 4,
		Value:  10,
	},
	}

	resStuff, resSum := solve(stuff, 10)
	fmt.Println(resStuff)
	fmt.Println(resSum)
}
