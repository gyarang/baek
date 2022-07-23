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

	var c string
	fmt.Fscanln(rd, &c)

	r := []rune(c)
	fmt.Fprintln(wr, int(r[0]))
}
