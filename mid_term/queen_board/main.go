package queen_board

//
//import (
//	"fmt"
//	"math"
//)
//
//var (
//	maxSum      int
//	boardValues [][]int
//)
//
//func main.go() {
//	boardValues = make([][]int, 8)
//	for i := 0; i < 8; i++ {
//		boardValues[i] = make([]int, 8)
//	}
//
//	var k int
//	fmt.Scan(&k)
//	for i := 0; i < k; i++ {
//		// Read board
//		for r := 0; r < 8; r++ {
//			for c := 0; c < 8; c++ {
//				fmt.Scan(&boardValues[r][c])
//			}
//		}
//
//		maxSum = math.MinInt
//		initialConfig := make([]int, 8)
//		solveNQueens(0, initialConfig)
//
//		fmt.Println(maxSum)
//	}
//
//}
//
//func solveNQueens(row int, currentConfig []int) {
//	if row == 8 {
//		currentSum := calculateSum(currentConfig)
//		if currentSum > maxSum {
//			maxSum = currentSum
//		}
//		return
//	}
//
//	for col := 0; col < 8; col++ {
//		if isValid(row, col, currentConfig) {
//			currentConfig[row] = col
//			solveNQueens(row+1, currentConfig)
//		}
//	}
//}
//
//func isValid(row, col int, currentConfig []int) bool {
//	for i := 0; i < row; i++ {
//		// Check for same column
//		if currentConfig[i] == col {
//			return false
//		}
//		// Check for diagonals (abs(col_diff) == abs(row_diff))
//		if abs(currentConfig[i]-col) == abs(i-row) {
//			return false
//		}
//	}
//	return true
//}
//
//func calculateSum(config []int) int {
//	sum := 0
//	for r := 0; r < 8; r++ {
//		c := config[r]
//		sum += boardValues[r][c]
//	}
//	return sum
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
