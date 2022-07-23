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

	var l, num int
	fmt.Fscanln(rd, &l, &num)

	var current int
	for i := 0; i < l; i++ {
		fmt.Fscanf(rd, "%d ", &current)
		if current < num {
			fmt.Fprintf(wr, "%d ", current)
		}
	}
	fmt.Fprintf(wr, "\n")
}
