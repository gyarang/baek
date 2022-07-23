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
	result := 1
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &input)
		result *= input
	}

	cnts := make(map[int]int)
	for result > 0 {
		remain := result % 10
		result = result / 10

		if _, ok := cnts[remain]; !ok {
			cnts[remain] = 1
		} else {
			cnts[remain] += 1
		}
	}

	for i := 0; i <= 9; i++ {
		if cnt, ok := cnts[i]; ok {
			fmt.Fprintln(writer, cnt)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}
