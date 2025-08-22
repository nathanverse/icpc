package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findFirstSmaller(dp []int, people [][]int, givenPerson []int) int {
	l, r := 0, len(dp)-1
	res := -1
	for l < r {
		mid := l + (r-l)/2
		midPerson := people[mid]
		if midPerson[0] < givenPerson[0] && midPerson[1] < givenPerson[1] {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}

func solve(n int, people [][]int) {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0]
	})

	dp := make([]int, 0)
	path := make([]int, n)

	for i := 0; i < n; i++ {
		// if a[i] > dp[len-1], append dp (+1), and update path
		if len(dp) == 0 {
			dp = append(dp, i)
			path[i] = -1
			continue
		}

		lastDpPerson := people[dp[len(dp)-1]]
		if people[i][0] > lastDpPerson[0] && people[i][1] > lastDpPerson[1] {
			dp = append(dp, i)
			path[i] = path[i-1]
			continue
		}

		firstSmallerPerson := findFirstSmaller(dp, people, people[i])
		if firstSmallerPerson == -1 {
			continue
		}
		path[i] = path[firstSmallerPerson]
		dp[firstSmallerPerson+1] = i
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintf(writer, "%d\n", len(dp))
	// Print path
	lastIndex := dp[len(dp)-1]
	for lastIndex >= 0 {
		if path[lastIndex] == -1 {
			fmt.Fprintf(writer, "%d", lastIndex+1)
		} else {
			fmt.Fprintf(writer, "%d ", lastIndex+1)
		}
		lastIndex = path[lastIndex]
	}

	return
}

func main() {
	// Create a new scanner to read from standard input.
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line, which is 'n'.
	scanner.Scan()
	nStr := scanner.Text()
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Error converting 'n' to integer:", err)
		return
	}

	// Initialize the 2D slice 'people'.
	var people [][]int

	// Loop 'n' times to read the subsequent lines.
	for i := 0; i < n; i++ {
		// Read a line.
		scanner.Scan()
		line := scanner.Text()

		// Split the line by spaces.
		parts := strings.Fields(line)

		// Create a new integer slice for the current person.
		person := make([]int, len(parts))

		// Convert string parts to integers.
		for j, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error converting part '%s' to integer: %v\n", part, err)
				return
			}
			person[j] = val
		}

		// Append the new integer slice to 'people'.
		people = append(people, person)
	}

	solve(n, people)
}
