package main

import (
	"bufio"
	"fmt"
	"os"
)

func pow(n, p int) int {
	if p == 0 {
		return 1
	}

	for i := 1; i < p; i++ {
		n *= n
	}
	return n
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b int
	fmt.Fscanf(reader, "%d %d", &a, &b)
	fmt.Fprint(writer, a*b)
}
