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

	var l int
	fmt.Fscanln(reader, &l)

	arr := make([]int, 10001)
	var num int
	for i := 0; i < l; i++ {
		fmt.Fscanf(reader, "%d\n", &num)
		arr[num] = arr[num] + 1
	}

	for i, v := range arr {
		for j := 0; j < v; j++ {
			fmt.Fprintln(writer, i)
		}
	}
}
