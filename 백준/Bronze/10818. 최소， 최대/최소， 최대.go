package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var l int
	fmt.Fscanln(rd, &l)

	var n int
	fmt.Fscanf(rd, "%d ", &n)
	min, max := n, n

	for i := 1; i < l; i++ {
		fmt.Fscanf(rd, "%d ", &n)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Fprintf(wr, "%d %d", min, max)
}
