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

	var input int
	fmt.Fscan(reader, &input)

	for input > 0 {
		fmt.Fprintln(writer, input)
		input -= 1
	}
}