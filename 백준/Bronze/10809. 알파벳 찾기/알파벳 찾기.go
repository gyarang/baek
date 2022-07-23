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

	var input string
	fmt.Fscan(reader, &input)

	dict := make(map[int32]int)
	for i, v := range input {
		if _, ok := dict[v]; !ok {
			dict[v] = i
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		if v, ok := dict[i]; ok {
			fmt.Fprintf(writer, "%d ", v)
		} else {
			fmt.Fprint(writer, "-1 ")
		}
	}
}
