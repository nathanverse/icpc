package main

import (
	"fmt"
	"slices"
	"sync"
	"unicode"
)

var moves = [][]int{{-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}}

func isVowel(c rune) bool {
	return slices.Contains([]rune{'a', 'e', 'y', 'u', 'o', 'i'}, unicode.ToLower(c))
}

func printRunes(runes []rune) string {
	s := ""
	for _, r := range runes {
		s += fmt.Sprintf("%c", r)
	}

	return s
}

// Constraint, a word should have exact 2 vowels, length of 4
func dfs(board *[][]rune, curX, curY int, curWord *[]rune, vowelCount int, res *map[string]struct{}) {
	backedBoard := *board
	backedCurWord := *curWord

	if len(backedCurWord) == 4 && vowelCount == 2 {
		if _, ok := (*res)[string(backedCurWord)]; !ok {
			(*res)[printRunes(backedCurWord)] = struct{}{}
		}
		return
	}

	for _, move := range moves {
		nextX := move[0] + curX
		nextY := move[1] + curY
		if nextX >= 0 && nextX < 4 && nextY >= 0 && nextY < 4 && backedBoard[nextX][nextY] != '.' {
			nextChar := backedBoard[nextX][nextY]

			if isVowel(nextChar) && vowelCount >= 2 {
				continue
			}

			vowelCountIndex := 0
			if isVowel(nextChar) {
				vowelCountIndex = 1
			}

			leftVowelCountToSeek := 2 - (vowelCount + vowelCountIndex)
			leftSlot := 4 - (len(backedCurWord) + 1)
			if leftSlot < leftVowelCountToSeek {
				continue
			}

			backedCurWord = append(backedCurWord, nextChar)
			backedBoard[nextX][nextY] = '.'
			dfs(&backedBoard, nextX, nextY, &backedCurWord, vowelCount+vowelCountIndex, res)
			backedBoard[nextX][nextY] = nextChar
			backedCurWord = backedCurWord[:len(backedCurWord)-1]
		}
	}
}

func solveBoggle(board1 [][]rune) map[string]struct{} {
	word := make([]rune, 0)
	resMap := map[string]struct{}{}
	for i := 0; i < len(board1); i++ {
		for j := 0; j < len(board1); j++ {
			curRune := board1[i][j]
			word = append(word, board1[i][j])
			board1[i][j] = '.'

			vowelCountIndex := 0
			if isVowel(curRune) {
				vowelCountIndex = 1
			}
			dfs(&board1, i, j, &word, vowelCountIndex, &resMap)
			board1[i][j] = curRune
			word = word[:len(word)-1]
		}
	}

	return resMap
}

func boggleGame(board1 [][]rune, board2 [][]rune) []string {
	var wg sync.WaitGroup
	var res1 map[string]struct{}
	var res2 map[string]struct{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		res1 = solveBoggle(board1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res2 = solveBoggle(board2)
	}()

	wg.Wait()

	ans := make([]string, 0)
	for k, _ := range res1 {
		if _, ok := res2[k]; ok {
			ans = append(ans, k)
		}
	}

	return ans
}

func main() {
	board1 := [][]rune{
		{'Z', 'W', 'A', 'V'},
		{'U', 'N', 'C', 'O'},
		{'Y', 'T', 'G', 'I'},
		{'H', 'G', 'P', 'M'},
	}
	board2 := [][]rune{
		{'G', 'S', 'F', 'U'},
		{'A', 'H', 'F', 'T'},
		{'G', 'N', 'A', 'L'},
		{'B', 'O', 'O', 'B'},
	}

	fmt.Println(boggleGame(board1, board2))
}
