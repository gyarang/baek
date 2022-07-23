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

	arr := make([]bool, 30)

	for i := 0; i < 28; i++ {
		var id int
		fmt.Fscanln(rd, &id)
		arr[id-1] = true
	}

	for i, v := range arr {
		if !v {
			fmt.Fprintln(wr, i+1)
		}
	}
}
