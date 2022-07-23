package main

import (
	"bufio"
	"fmt"
	"os"
)

var rd = bufio.NewReader(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func checkAsc(input []int) {
	for i := 1; i < 9; i++ {
		if input[i-1] != i {
			fmt.Fprint(wr, "mixed")
			return
		}
	}
	fmt.Fprint(wr, "ascending")
}

func checkDesc(input []int) {
	for i := 0; i < 8; i++ {
		if input[i] != 8-i {
			fmt.Fprint(wr, "mixed")
			return
		}
	}
	fmt.Fprint(wr, "descending")
}

func main() {
	defer wr.Flush()

	arr := make([]int, 8)
	fmt.Fscanf(rd, "%d ", &arr[0])
	if arr[0] != 1 && arr[0] != 8 {
		fmt.Fprint(wr, "mixed")
		return
	}

	for i := 1; i < 8; i++ {
		fmt.Fscanf(rd, "%d ", &arr[i])
	}

	if arr[0] == 1 {
		checkAsc(arr)
	} else {
		checkDesc(arr)
	}
}
