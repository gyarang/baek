package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cnt int
	fmt.Fscanln(reader, &cnt)
	times := make([]int, cnt, cnt)

	for i := 0; i < cnt; i++ {
		fmt.Fscan(reader, &times[i])
	}
	sort.Ints(times)

	max, temp := 0, 0
	for i := 0; i < cnt; i++ {
		temp = temp + times[i]
		max += temp
	}

	fmt.Fprintln(writer, max)
}