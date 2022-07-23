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

	fmt.Fscan(reader, &n)

	switch {
	case n >= 90:
		fmt.Fprint(writer, "A")
	case n >= 80:
		fmt.Fprint(writer, "B")
	case n >= 70:
		fmt.Fprint(writer, "C")
	case n >= 60:
		fmt.Fprint(writer, "D")
	default:
		fmt.Fprint(writer, "F")
	}
}
