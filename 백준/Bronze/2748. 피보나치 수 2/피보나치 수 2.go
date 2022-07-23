package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibonacci(cnt int) int {
	var result int
	for i, first, second := 0, 0, 1; i <= cnt; i, first, second = i+1, first+second, first {
		if i == cnt {
			result = first
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	result := fibonacci(n)
	fmt.Fprintln(writer, result)
}
