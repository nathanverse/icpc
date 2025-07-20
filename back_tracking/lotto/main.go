package main

import "fmt"

func Solve(nums []int, l, r, from int) {
	if r < 6 {
		fmt.Print(-1)
	}

	if l == 6 {
		for i := 0; i < 6; i++ {
			fmt.Print(nums[i])

			if i != 5 {
				fmt.Print(" ")
			}
		}

		fmt.Println()

		return
	}

	for i := from; i < r; i++ {
		nums[i], nums[l] = nums[l], nums[i]
		Solve(nums, l+1, r, i+1)
		nums[i], nums[l] = nums[l], nums[i]
	}
}

func main() {
	n := -1

	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			panic(err)
		}

		if n == 0 {
			break
		}

		nums := make([]int, n)
		for i := 0; i < n; i++ {
			_, err := fmt.Scan(&nums[i])
			if err != nil {
				panic(err)
			}
		}

		Solve(nums, 0, n, 0)

		fmt.Println()
	}
}
