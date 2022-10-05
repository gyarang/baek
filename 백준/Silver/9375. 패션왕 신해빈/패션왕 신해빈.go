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

	tcnt := 0
	fmt.Fscanln(reader, &tcnt)

	for i := 0; i < tcnt; i++ {
		var cnt int
		fmt.Fscanln(reader, &cnt)
		counter := make(map[string]int)
		for j := 0; j < cnt; j++ {
			var cloth string
			fmt.Fscanln(reader, &cloth, &cloth)
			counter[cloth]++
		}

		result := 1
		for k := range counter {
			result *= counter[k] + 1
		}
		fmt.Fprintln(writer, result - 1)
	}
}
