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

	var len int
	fmt.Fscanln(reader, &len)

	inputs := make([]int, len)
	var max int
	for i := 0; i < len; i++ {
		fmt.Fscanf(reader, "%d ", &inputs[i])
		if inputs[i] > max {
			max = inputs[i]
		}
	}

	var sum float32
	for _, v := range inputs {
		newPoint := float32(v) / float32(max) * 100.0
		sum += newPoint
	}

	fmt.Fprintln(writer, sum/float32(len))
}
