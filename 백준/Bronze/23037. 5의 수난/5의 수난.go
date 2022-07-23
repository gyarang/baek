package main

import (
	"bufio"
	"fmt"
	"os"
)

func power(num, cnt int) int {
	result := 1
	for i := 0; i < cnt; i++ {
		result *= num
	}
	return result
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var sum, n int
	for i := 0; i <= 4; i++ {
		fmt.Fscanf(rd, "%1d", &n)
		sum += power(n, 5)
	}
	fmt.Fprintln(wr, sum)
}
