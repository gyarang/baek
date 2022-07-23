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

	var cnt int
	fmt.Fscanln(rd, &cnt)

	var result, num int
	for i := 0; i < cnt; i++ {
		fmt.Fscanf(rd, "%1d", &num)
		result += num
	}

	fmt.Fprintln(wr, result)
}
