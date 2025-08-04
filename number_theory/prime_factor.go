package main

import (
	"errors"
	"fmt"
)

func solve(n int) []int {
	if n == 1 {
		return []int{1}
	}

	if n <= 0 {
		panic(errors.New("n must be > 0"))
	}

	res := make([]int, 0)
	for n%2 == 0 {
		res = append(res, 2)
		n /= 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			res = append(res, i)
			n /= i
		}
	}

	if n > 2 {
		res = append(res, n)
	}

	return res
}

func main() {
	fmt.Println(solve(252)) // <- expect:2,2,3,3,7
}
