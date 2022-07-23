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

	var cnt int
	fmt.Fscanln(rd, &cnt)

	arr := make([][]uint8, cnt)
	var str string
	for i := 0; i < cnt; i++ {
		fmt.Fscanln(rd, &str)
		arr[i] = []uint8(str)
	}

	inputLength := len(arr[0])

L1:
	for i := 0; i < inputLength; i++ {
		for _, s := range arr[1:] {
			if s[i] != arr[0][i] {
				fmt.Fprint(wr, "?")
				continue L1
			}
		}
		fmt.Fprint(wr, string(arr[0][i]))
	}
}
