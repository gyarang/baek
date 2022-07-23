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

	var n, idx, max int
	for i := 0; i < 9; i++ {
		fmt.Fscanf(rd, "%d\n", &n)
		if n > max {
			idx = i + 1
			max = n
		}
	}

	fmt.Fprintf(wr, "%d\n%d", max, idx)
}
