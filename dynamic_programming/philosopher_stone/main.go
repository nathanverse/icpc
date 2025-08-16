package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxx(x, y int) int {
	maxxx := x
	if y > x {
		maxxx = y
	}

	return maxxx
}

func readInput(reader *bufio.Scanner) (h, w int, stones [][]int) {

	// Read h and w
	reader.Scan()
	h, _ = strconv.Atoi(reader.Text())
	reader.Scan()
	w, _ = strconv.Atoi(reader.Text())

	// Read the stone grid
	stones = make([][]int, h)
	for i := 0; i < h; i++ {
		stones[i] = make([]int, w)
		for j := 0; j < w; j++ {
			reader.Scan()
			stones[i][j], _ = strconv.Atoi(reader.Text())
		}
	}
	return h, w, stones
}

func solve(h, w int, stones [][]int) int64 {
	if h < 0 || w < 0 {
		panic("h < w")
	}

	if h > 1 {
		for i := 1; i < h; i++ {
			last2Max := stones[i-1][0]
			for j := 0; j < w; j++ {
				if j == w-1 {
					stones[i][j] = last2Max + stones[i][j]
					continue
				}
				stones[i][j] = maxx(last2Max+stones[i][j], stones[i][j]+stones[i-1][j+1])
				last2Max = maxx(stones[i-1][j], stones[i-1][j+1])
			}
		}
	}

	res := -1
	for i := 0; i < w; i++ {
		res = maxx(res, stones[h-1][i])
	}

	return int64(res)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanWords)
	reader.Scan()
	T, _ := strconv.Atoi(reader.Text())

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for t := 0; t < T; t++ {
		h, w, stones := readInput(reader)
		writer.WriteString(fmt.Sprintf("%d\n", solve(h, w, stones)))
	}
}
