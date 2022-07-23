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

	var h, m int
	fmt.Fscanf(reader, "%d %d", &h, &m)

	m -= 45
	if m < 0 {
		m += 60
		h -= 1
		if h < 0 {
			h = 23
		}
	}

	fmt.Fprintf(writer, "%d %d", h, m)
}
