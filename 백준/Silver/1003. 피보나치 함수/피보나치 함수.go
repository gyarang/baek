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

	var cnt, input int
	fmt.Fscanln(reader, &cnt)

	arr := make([]int, 42)
	arr[1] = 1
	arr[2] = 1

	max := 2

	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &input)
		if input == 0 {
			fmt.Fprintln(writer, "1 0")
			continue
		}
		if input > max {
			for j := max + 1; j <= input; j++ {
				arr[j] = arr[j-1] + arr[j-2]
			}
			max = input
		}
		fmt.Fprintln(writer, arr[input-1], arr[input])
	}
}
