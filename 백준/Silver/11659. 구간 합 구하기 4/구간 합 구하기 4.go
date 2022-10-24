package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, M int
	fmt.Fscanln(reader, &N, &M)
	arr := [100001]int{}

	fmt.Fscan(reader, &arr[1])
	for i := 2; i < N+1; i++ {
		var value int
		fmt.Fscan(reader, &value)
		arr[i] = value + arr[i-1]
	}
	reader.ReadBytes('\n')

	for i := 0; i < M; i++ {
		var start, end int
		fmt.Fscanln(reader, &start, &end)
		fmt.Fprintln(writer, arr[end]-arr[start-1])
	}
}
