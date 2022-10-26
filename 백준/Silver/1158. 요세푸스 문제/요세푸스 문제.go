package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, K int
	fmt.Scanln(&N, &K)

	r := ring.New(N)
	for i := 1; i <= N; i++ {
		r.Value = i
		r = r.Next()
	}
	r = r.Prev()

	fmt.Fprint(writer, "<")
	for r.Len() > 1 {
		r = r.Move(K)
		fmt.Fprintf(writer, "%d, ", r.Value)
		r = r.Prev()
		r.Unlink(1)
	}
	fmt.Fprintf(writer, "%d>\n", r.Value)
}
