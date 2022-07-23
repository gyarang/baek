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

	var a, b int
	for {
		cnt, _ := fmt.Fscanln(rd, &a, &b)
		if cnt != 2 {
			break
		}
		fmt.Fprintln(wr, a+b)
	}
}
