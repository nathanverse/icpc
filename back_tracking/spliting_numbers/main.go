package main

import "fmt"

func findAAndB(n int) {
	curBit := 0
	totalOfBit1 := 0
	a, b := 0, 0
	for n > 0 {
		if n&1 == 0 {
			curBit++
			n >>= 1
			continue
		}

		totalOfBit1 += 1
		if totalOfBit1%2 == 1 {
			a = a | 1<<curBit
		} else {
			b = b | 1<<curBit
		}
		curBit++
		n >>= 1
	}

	fmt.Printf("%d %d\n", a, b)
}

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}

		if n == 0 {
			break
		}

		findAAndB(n)
	}
}
