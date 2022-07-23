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

	var x, y, w, h int
	fmt.Fscanf(reader, "%d %d %d %d", &x, &y, &w, &h)

	if x > w/2 {
		x = w - x
	}
	if y > h/2 {
		y = h - y
	}

	if x >= y {
		fmt.Fprint(writer, y)
	} else {
		fmt.Fprint(writer, x)
	}
}
