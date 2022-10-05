package main

import "fmt"

func main() {
	dp := make([]int, 1000001)
	dp[0], dp[1], dp[2], dp[3] = 0, 0, 1, 1

	var input int
	fmt.Scan(&input)

	if input < 3 {
		fmt.Println(dp[input])
		return
	}

	for i := 4; i <= input; i++ {
		m1 := dp[i-1]
		d2 := dp[i/2]
		d3 := dp[i/3]

		if i % 2 != 0 && i % 3 != 0 {
			dp[i] = m1 + 1
			continue
		} else if i % 2 == 0 && i % 3 != 0 {
			dp[i] = min(m1, d2) + 1
			continue
		} else if i % 2 != 0 && i % 3 == 0 {
			dp[i] = min(m1, d3) + 1
			continue
		}

		dp[i] = min(m1, d2, d3) + 1
	}

	fmt.Println(dp[input])
}

func min(input ...int) int {
	min := input[0]
	for i := 1; i < len(input); i++ {
		if min > input[i] {
			min = input[i]
		}
	}
	return min
}
