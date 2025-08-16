package main

import "fmt"

func change(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}

	combArray := make([]int, amount+1)
	combArray[0] = 1
	for i := 0; i < len(coins); i++ {
		curCoin := coins[i]
		for curAmount := curCoin; curAmount <= amount; curAmount++ {
			combArray[curAmount] += combArray[curAmount-curCoin]
		}
	}

	return combArray[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
