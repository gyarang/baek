package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	now := 1
	mov := 1
	increase := 6

	for now < n {
		mov += 1
		now += increase
		increase += 6
	}

	fmt.Fprintln(writer, mov)
}
