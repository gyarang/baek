package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Fprint(wr, " ")
		}

		for j := 0; j < i; j++ {
			fmt.Fprint(wr, "*")
		}

		fmt.Fprint(wr, "\n")
	}
}
