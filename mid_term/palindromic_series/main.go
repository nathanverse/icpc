package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	baseDigits := make([]int, 0, len(numStr))
	totalSumOfDigits := 0

	for _, charDigit := range numStr {
		digit := int(charDigit - '0')
		baseDigits = append(baseDigits, digit)
		totalSumOfDigits += digit
	}

	if totalSumOfDigits <= 1 {
		return true
	}

	lenP := len(baseDigits)
	for i := 0; i < totalSumOfDigits/2; i++ {
		digitLeft := baseDigits[i%lenP]
		digitRight := baseDigits[(totalSumOfDigits-1-i)%lenP]

		if digitLeft != digitRight {
			return false
		}
	}

	return true
}

func main() {
	var T int
	fmt.Scan(&T) // Read the number of test cases

	for i := 0; i < T; i++ {
		var num int
		fmt.Scan(&num) // Read the number N for each test case

		if IsPalindrome(num) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
