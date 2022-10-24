package main

import "fmt"

func main() {
	var width int
	fmt.Scanln(&width)

	dp := [1001]int{}
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= width; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= 10007
	}

	fmt.Println(dp[width])
}
