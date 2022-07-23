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
	fmt.Fscanf(rd, "%d %d\n", &a, &b)

	fmt.Fprintf(wr, "%d\n%d\n%d\n%d\n%d", a+b, a-b, a*b, a/b, a%b)
}
