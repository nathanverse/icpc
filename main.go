package main

import "fmt"

// smallPowerMod calculates (a^b) % m using binary exponentiation (a.k.a. exponentiation by squaring).
// This is an efficient way to compute large powers.
//func smallPowerMod(a, b, m int) int {
//	res := 1
//	a %= m
//	for b > 0 {
//		// If b is odd, multiply a with the result.
//		if b%2 == 1 {
//			res = (res * a) % m
//		}
//		// b must be even now, so we can square a and half b.
//		a = (a * a) % m
//		b /= 2
//	}
//	return res
//}
//
//func superPowerMod(a int, b []int) int {
//	if a <= 1 {
//		return a
//	}
//
//	if len(b) == 0 {
//		return 1
//	}
//
//	mod := 1337
//	powerOf10Base := smallPowerMod(a, 10, mod)
//
//	res := 1
//	if b[len(b)-1]%10 != 0 {
//		res = smallPowerMod(a, b[len(b)-1], mod)
//	}
//
//	for i := len(b) - 2; i >= 0; i-- {
//		digit := b[i]
//
//		if digit%10 != 0 {
//			res = (smallPowerMod(powerOf10Base, digit, mod) * res) % mod
//		}
//		// a^100 % m = ((a^10) % m)^10 % m
//		powerOf10Base = smallPowerMod(powerOf10Base, 10, mod)
//	}
//	return res
//}
//
//func factorial(n int) int {
//	if n == 0 || n == 1 {
//		return 1
//	}
//	return n * factorial(n-1)
//}
//
//// The boggle game
//func probabilityToDestination(originalSigns, detectedSigns []rune) float64 {
//	minusCount, plusCount, totalMissingCount := 0, 0, 0
//	for i, v := range detectedSigns {
//		if v == '?' {
//			if originalSigns[i] == '+' {
//				plusCount++
//			} else {
//				minusCount++
//			}
//
//			totalMissingCount++
//		}
//
//		continue
//	}
//
//	totalCombinations := math.Pow(2, float64(totalMissingCount))
//	possibleCombinations := factorial(totalMissingCount) / (factorial(minusCount) * factorial(plusCount))
//
//	return float64(possibleCombinations) / totalCombinations
//}

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//type LetterState struct {
//	IsDone    bool
//	IsCapital bool
//}
//
//func Solution(letters string) int {
//	// Implement your solution here
//
//	letterStateMap := map[rune]*LetterState{}
//	res := 0
//	for _, v := range letters {
//		if letterState, ok := letterStateMap[unicode.ToLower(v)]; ok {
//			if letterState.IsDone {
//				continue
//			}
//
//			if v == unicode.ToLower(v) && letterState.IsCapital {
//				res--
//				letterState.IsDone = true
//			}
//
//			if !letterState.IsCapital && v != unicode.ToLower(v) {
//				letterStateMap[unicode.ToLower(v)].IsCapital = true
//				res++
//			}
//
//			continue
//		} else {
//			if unicode.ToLower(v) != v {
//				letterStateMap[unicode.ToLower(v)] = &LetterState{
//					IsDone:    true,
//					IsCapital: false,
//				}
//				continue
//			}
//
//			letterStateMap[v] = &LetterState{
//				IsDone:    false,
//				IsCapital: false,
//			}
//		}
//	}
//
//	return res
//}

//func maxSumAfterPartitioning(arr []int, k int) int {
//	maxSum := -1
//	maxSumAfterPartitioningRecursion(arr, k, 0, 0, &maxSum)
//
//	return maxSum
//}
//
//func maxSumAfterPartitioningRecursion(arr []int, k int, start int, state int, maxSum *int) {
//	if start == len(arr) {
//		*maxSum = max(*maxSum, state)
//
//		return
//	}
//
//	count := 0
//	maxEle := -1
//	for i := start; i < len(arr) && count < k; i++ {
//		maxEle = max(maxEle, arr[i])
//		curSum := maxEle * (i - start + 1)
//		maxSumAfterPartitioningRecursion(arr, k, i+1, state+curSum, maxSum)
//		count++
//	}
//
//	return
//}

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr)+1)
	curMax := -1
	for j := len(arr) - 1; j >= 0; j-- {
		if len(arr)-j <= k {
			curMax = max(arr[j], curMax)
			dp[j] = curMax * (len(arr) - j)
			continue
		}

		curMax = arr[j]
		dpJ := arr[j] + dp[j+1]
		for i := 1; i < k; i++ {
			curMax = max(curMax, arr[j+i])
			maxDpI := curMax * (i + 1)
			dpJ = max(maxDpI+dp[j+i+1], dpJ)
		}

		dp[j] = dpJ
	}

	return dp[0]
}

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
}
