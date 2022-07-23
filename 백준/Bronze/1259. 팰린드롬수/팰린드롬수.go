package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := -1
	arr := make([]int, 0, 100)
	for n != 0 {
		fmt.Fscan(reader, &n)
		if n != 0 {
			arr = append(arr, n)
		}
	}

L1:
	for _, v := range arr {
		var numArr []int
		for v > 0 {
			numArr = append(numArr, v%10)
			v = v / 10
		}
		for i := 0; i < len(numArr)/2; i++ {
			if numArr[i] != numArr[len(numArr)-1-i] {
				fmt.Fprintln(writer, "no")
				continue L1
			}
		}
		fmt.Fprintln(writer, "yes")
	}
}
