package main

import (
	"errors"
	"fmt"
)

// Given an integer n -> Return all prime numbers from 1 to n (inclusive)
func sieveOfEratos(n int) []int {
	if n <= 0 {
		panic(errors.New("n must be > 0"))
	}

	isNotPrime := make([]bool, n+1)
	isNotPrime[0] = true
	isNotPrime[1] = true
	res := make([]int, 0)

	i := 2
	for i*i <= n {
		if !isNotPrime[i] {
			res = append(res, i)
			for j := i * i; j <= n; j += i {
				isNotPrime[j] = true
			}
		}

		i++
	}

	for i <= n {
		if !isNotPrime[i] {
			res = append(res, i)
		}
		i++
	}

	return res
}

func main() {
	fmt.Println(sieveOfEratos(40))

	//2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37.
}
