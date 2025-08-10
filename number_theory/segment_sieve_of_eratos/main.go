package main

import (
	"errors"
	"fmt"
	"math"
)

const BLOCK_SIZE = 3

// segmentSieve Given l, r are integers, return all prime numbers between them.
// With 1 < l < r < 10^16.

// With naive approach, we need to store O(N), which mean 10^16 bytes -> 1TB
// With memory optimized approach, only O(T), with T is the number of primes from 1 -> sqrt(N).
// 10^8 has around 6mil number of primes, which means we need to store 6*4 mil bytes, approximately 24MB.

func regularSieve(n int) []int {
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

// Inclusive
func segmentSieve(l, r int) []int {
	sqrtR := int(math.Sqrt(float64(r)))
	baseSegments := regularSieve(sqrtR)
	res := make([]int, 0)
	curBlockIndex := l
	for curBlockIndex <= r {
		curBlockEndIndex := curBlockIndex + BLOCK_SIZE
		if curBlockEndIndex >= r {
			curBlockEndIndex = r
		}

		isNotPrimes := make([]bool, curBlockEndIndex-curBlockIndex+1)
		for _, v := range baseSegments {
			if v*v > curBlockEndIndex {
				break
			}

			firstMutiplier := curBlockIndex / v
			if firstMutiplier*v < curBlockIndex {
				firstMutiplier++
			}

			if firstMutiplier <= 1 {
				firstMutiplier = 2
			}

			for firstMutiplier*v <= curBlockEndIndex {
				standardizedIndex := firstMutiplier*v - curBlockIndex
				isNotPrimes[standardizedIndex] = true
				firstMutiplier++
			}
		}

		for i, isNotPrime := range isNotPrimes {
			if !isNotPrime && i+curBlockIndex > 1 {
				res = append(res, i+curBlockIndex)
			}
		}

		curBlockIndex += BLOCK_SIZE
	}

	return res
}

func main() {
	fmt.Println(segmentSieve(10, 50))
}
