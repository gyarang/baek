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

	cnt := 0
	fmt.Fscanln(reader, &cnt)
	arr := make([]int, cnt)
	val := 0
	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &val)
		arr[i] = val
	}

	sort.Ints(arr)
	for _, v := range arr {
		fmt.Fprintln(writer, v)
	}
}
