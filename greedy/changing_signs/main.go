package main

import "fmt"

func solve(nums []int, changes int) int {
	if changes < 0 {
		panic("changes < 0")
	}

	k := changes
	res := 0
	metPositive := false
	for i := 0; i < len(nums); i++ {
		if k == 0 {
			res += nums[i]
			continue
		}

		if nums[i] < 0 {
			res -= nums[i]
			k--
			continue
		}

		if !metPositive {
			metPositive = true
			m := k
			k = 0
			if m > 0 && m%2 == 1 {
				if i == 0 || -nums[i-1] > nums[i] {
					res -= nums[i]
					continue
				}

				res = res + nums[i-1]*2 + nums[i]
				continue
			}
			res += nums[i]
			continue
		}

		res += nums[i]
	}

	if !metPositive && k > 0 && k%2 == 1 {
		res += 2 * nums[len(nums)-1]
	}

	return res
}

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		panic(err)
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		_, err2 := fmt.Scan(&nums[i])
		if err2 != nil {
			panic(err)
		}
	}

	fmt.Println(solve(nums, k))
}
