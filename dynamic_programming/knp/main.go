package main

import "fmt"

func execute(s1, s2 string) []int {
	if len(s1) < 1 || len(s2) < 1 {
		panic("<UNK>")
	}

	if len(s1) == len(s2) {
		if s1 == s2 {
			return []int{0}
		}

		return []int{}
	}

	// make s1 smaller string
	if len(s2) < len(s1) {
		s1, s2 = s2, s1
	}

	j := 0
	prefix := make([]int, len(s1))
	for i := 0; i < len(s1); i++ {
		if s1[i] == s1[j] && i != 0 {
			prefix[i] = j
			j++
			continue
		}

		j = 0
	}

	j = -1
	i := -1
	ans := make([]int, 0)
	for i < len(s2)-1 {
		if j < len(s1)-1 && s2[i+1] == s1[j+1] {
			if j+1 == len(s1)-1 {
				ans = append(ans, i+2-len(s1))
			}

			j++
			i++
			continue
		}

		if j > 0 {
			j = prefix[j]
			continue
		}

		i++
	}

	return ans
}

func main() {
	s1 := "ABABDABACDABABCABAB"
	fmt.Println(execute(s1, "ABABCABAB"))
}
