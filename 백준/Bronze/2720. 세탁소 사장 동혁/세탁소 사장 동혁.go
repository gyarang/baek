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

	var cnt, remain int
	fmt.Fscanln(reader, &cnt)

	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &remain)

		var q, d, n, p int
		q = remain / 25
		remain %= 25
		d = remain / 10
		remain %= 10
		n = remain / 5
		remain %= 5
		p = remain

		fmt.Fprintln(writer, q, d, n, p)
	}
}
