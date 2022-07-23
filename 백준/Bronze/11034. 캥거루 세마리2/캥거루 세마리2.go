package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b, c int) int {
	if b-a > c-b {
		return b - a
	} else {
		return c - b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var a, b, c int
		cnt, _ := fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		if cnt != 3 {
			return
		}
		m := max(a, b, c)
		fmt.Fprintln(writer, m-1)
	}
}
