package main

import (
	"bufio"
	"fmt"
	"os"
)

var apart [15][15]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < 15; i++ {
		apart[0][i] = i
	}

	for i := 1; i < 15; i++ {
		apart[i][0] = apart[i-1][0]
		for j := 1; j < 15; j++ {
			apart[i][j] = apart[i][j-1] + apart[i-1][j]
		}
	}

	var cnt, k, n int
	fmt.Fscanln(reader, &cnt)
	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n)
		fmt.Fprintln(writer, apart[k][n])
	}
}
