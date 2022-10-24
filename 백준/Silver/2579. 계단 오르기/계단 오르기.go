package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cnt int
	fmt.Fscanln(reader, &cnt)
	stairs := [301]int{}
	dp := [301]int{}

	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &stairs[i])
	}

	dp[0] = stairs[0]
	dp[1] = max(stairs[0]+stairs[1], stairs[0])
	dp[2] = max(stairs[0]+stairs[2], stairs[1]+stairs[2])

	for i := 3; i < cnt; i++ {
		dp[i] = max(dp[i-2]+stairs[i], stairs[i-1]+stairs[i]+dp[i-3])
	}

	fmt.Fprintln(writer, dp[cnt-1])
}
