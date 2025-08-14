package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Using polynomial hash helps us to reduce
// the complexity of the algo into O(n) from O(n^2) as we don't need to copy
// the suffix
func suffixPrefix(str string) int {
	base := 31
	m := int(1e9 + 7)

	prefixHash := 0
	suffixHash := 0
	power := 1
	n := len(str)
	res := 0
	for i := 0; i < n-1; i++ {
		prefixChar := int(str[i] - 'a' + 1)
		prefixHash = (prefixHash + prefixChar*power) % m
		power = (power * base) % m

		suffixChar := int(str[n-i-1] - 'a' + 1)
		suffixHash = (suffixChar + suffixHash*base) % m

		if prefixHash == suffixHash {
			res++
		}
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	nWord := scanner.Text()

	n, err := strconv.Atoi(nWord)
	if err != nil {
		fmt.Println("Error converting input to an integer:", err)
		return
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := 0; i < n; i++ {
		scanner.Scan()
		nextInput := scanner.Text()

		fmt.Fprintf(writer, "Case %d: %d\n", i+1, suffixPrefix(nextInput))
	}
}
