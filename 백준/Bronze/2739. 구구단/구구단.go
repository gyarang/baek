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

	var a int
	fmt.Fscanf(rd, "%d", &a)
	for i := 1; i < 10; i++ {
		fmt.Fprintf(wr, "%d * %d = %d\n", a, i, a*i)
	}
}
