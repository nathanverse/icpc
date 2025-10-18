package main

import (
	"math"
)

// smallPowerMod calculates (a^b) % m using binary exponentiation (a.k.a. exponentiation by squaring).
// This is an efficient way to compute large powers.
func smallPowerMod(a, b, m int) int {
	res := 1
	a %= m
	for b > 0 {
		// If b is odd, multiply a with the result.
		if b%2 == 1 {
			res = (res * a) % m
		}
		// b must be even now, so we can square a and half b.
		a = (a * a) % m
		b /= 2
	}
	return res
}

func superPowerMod(a int, b []int) int {
	if a <= 1 {
		return a
	}

	if len(b) == 0 {
		return 1
	}

	mod := 1337
	powerOf10Base := smallPowerMod(a, 10, mod)

	res := 1
	if b[len(b)-1]%10 != 0 {
		res = smallPowerMod(a, b[len(b)-1], mod)
	}

	for i := len(b) - 2; i >= 0; i-- {
		digit := b[i]

		if digit%10 != 0 {
			res = (smallPowerMod(powerOf10Base, digit, mod) * res) % mod
		}
		// a^100 % m = ((a^10) % m)^10 % m
		powerOf10Base = smallPowerMod(powerOf10Base, 10, mod)
	}
	return res
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// The boggle game
func probabilityToDestination(originalSigns, detectedSigns []rune) float64 {
	minusCount, plusCount, totalMissingCount := 0, 0, 0
	for i, v := range detectedSigns {
		if v == '?' {
			if originalSigns[i] == '+' {
				plusCount++
			} else {
				minusCount++
			}

			totalMissingCount++
		}

		continue
	}

	totalCombinations := math.Pow(2, float64(totalMissingCount))
	possibleCombinations := factorial(totalMissingCount) / (factorial(minusCount) * factorial(plusCount))

	return float64(possibleCombinations) / totalCombinations
}

func main() {
	//a := 2
	//b := []int{1, 0} // Represents the number 123
	//
	//result := superPowerMod(a, b)
	//fmt.Printf("%d\n", result) // Correct output is 2^123 mod 1337

	probabilityToDestination([]rune{'+', '-', '+', '-'}, []rune{'+', '-', '?', '?'})
}
