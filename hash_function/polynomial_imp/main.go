package main

import (
	"bufio"
	"fmt"
	"os"
)

func polynomial_hash(s string) int {
	base := 31
	m := int(1e9 + 7) // large prime number
	hashValue := 0
	power := 1
	for i := 0; i < len(s); i++ {
		intHashValue := int(s[i] - 'a' + 1)
		hashValue = (hashValue + intHashValue*power) % m
		power = (power * base) % m
	}

	return hashValue
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintf(writer, "%d\n", polynomial_hash("geeksforgeeks"))
}
