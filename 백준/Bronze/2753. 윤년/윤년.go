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

	if input%4 == 0 {
		if input%400 == 0 {
			fmt.Fprint(writer, "1")
			return
		} else if input%100 == 0 {
			fmt.Fprint(writer, "0")
			return
		} else {
			fmt.Fprint(writer, "1")
			return
		}
	}
	fmt.Fprint(writer, "0")
}
