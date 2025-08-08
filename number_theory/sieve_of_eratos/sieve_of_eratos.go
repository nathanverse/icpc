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
	int64n := int64(n)

	for i := 2; i <= n; i++ {
		if int64(i*i) > int64n {
			break
		}

		if !isNotPrime[i] {
			for j := i * i; j <= n; j += i {
				isNotPrime[j] = true
			}
		}
	}

	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		if !isNotPrime[i] {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	fmt.Println(sieveOfEratos(40))

	//2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37.
}
