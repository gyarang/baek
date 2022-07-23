package main

import (
	"bufio"
	"fmt"
	"os"
)

type body struct {
	x, y, rank int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cnt int
	fmt.Fscanln(reader, &cnt)

	bodies := make([]body, cnt)
	for i := 0; i < cnt; i++ {
		b := body{rank: 1}
		fmt.Fscanln(reader, &b.y, &b.x)
		bodies[i] = b
	}

	for i := 0; i < cnt; i++ {
		for j := 0; j < cnt; j++ {
			if bodies[i].x < bodies[j].x && bodies[i].y < bodies[j].y {
				bodies[i].rank++
			}
		}
	}

	for _, v := range bodies {
		fmt.Fprintf(writer, "%d ", v.rank)
	}
}
