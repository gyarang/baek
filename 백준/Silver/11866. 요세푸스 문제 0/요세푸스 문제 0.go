package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, cnt int
	fmt.Fscanf(reader, "%d %d", &l, &cnt)

	outs := make([]int, 0, l)
	r := ring.New(l)
	for i := 0; i < l; i++ {
		r.Value = i + 1
		r = r.Next()
	}

	r = r.Move(-1)
	for len(outs) != l {
		r = r.Move(cnt - 1)
		out := r.Unlink(1)
		outs = append(outs, out.Value.(int))
	}

	fmt.Fprint(writer, "<")
	for i, v := range outs {
		fmt.Fprint(writer, v)
		if i < len(outs)-1 {
			fmt.Fprint(writer, ", ")
		}
	}
	fmt.Fprint(writer, ">")
}
