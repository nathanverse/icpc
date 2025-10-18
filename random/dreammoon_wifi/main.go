package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func probabilityToDestination(originalSigns, detectedSigns []rune) float64 {
	orgMinusCount, orgPlusCount := 0, 0
	detectedMinusCount, detectedPlusCount := 0, 0
	missingCount := 0
	for i := 0; i < len(originalSigns); i++ {
		if originalSigns[i] == '-' {
			orgMinusCount++
		} else {
			orgPlusCount++
		}

		if detectedSigns[i] == '-' {
			detectedMinusCount++
		} else if detectedSigns[i] == '+' {
			detectedPlusCount++
		} else {
			missingCount++
		}
	}

	leftMinusToFill := orgMinusCount - detectedMinusCount
	leftPlusToFill := orgPlusCount - detectedPlusCount
	if leftMinusToFill < 0 || leftPlusToFill < 0 {
		return 0
	}

	totalCombinations := math.Pow(2, float64(missingCount))
	possibleCombinations := factorial(missingCount) / (factorial(leftMinusToFill) * factorial(leftPlusToFill))

	if totalCombinations == 0 {
		panic("No combinations found")
	}

	return float64(possibleCombinations) / totalCombinations
}

func main() {
	runes1 := make([]rune, 0)
	runes2 := make([]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		runes1 = []rune(scanner.Text())
	}

	if scanner.Scan() {
		runes2 = []rune(scanner.Text())
	}

	fmt.Println(probabilityToDestination(runes1, runes2))
}
