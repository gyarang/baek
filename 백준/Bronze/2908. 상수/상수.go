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

	var a, b, ra, rb int

	fmt.Fscanf(reader, "%d %d", &a, &b)

	for i := 2; i >= 0; i-- {
		p := pow(10, i)
		ra += (a % 10) * p
		rb += (b % 10) * p
		a /= 10
		b /= 10
	}

	if ra > rb {
		fmt.Fprint(writer, ra)
	} else {
		fmt.Fprint(writer, rb)
	}
}
