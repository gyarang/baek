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

	var n int
	fmt.Fscan(rd, &n)

	result := 1
	for i := 2; i <= n; i++ {
		result = result * i
	}

	fmt.Fprintf(wr, "%d\n", result)
}
