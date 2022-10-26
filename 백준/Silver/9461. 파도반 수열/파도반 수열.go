package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cnt int
	dp := make([]int, 101)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 1
	last := 3

	fmt.Fscanln(reader, &cnt)
	for i := 0; i < cnt; i++ {
		var input int
		fmt.Fscanln(reader, &input)
		if input <= last {
			fmt.Fprintln(writer, dp[input])
			continue
		}
		for j := last+1; j <= input; j++ {
			dp[j] = dp[j-2] + dp[j-3]
		}
		last = input
		fmt.Fprintln(writer, dp[input])
	}
}
