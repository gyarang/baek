package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var l int
	fmt.Fscanln(rd, &l)

	arr := make([]int, l)
	for i := range arr {
		fmt.Fscanf(rd, "%d ", &arr[i])
	}

	var num int
	fmt.Fscan(rd, &num)

	cnt := 0
	for _, v := range arr {
		if v == num {
			cnt += 1
		}
	}

	fmt.Fprintln(wr, cnt)
}
