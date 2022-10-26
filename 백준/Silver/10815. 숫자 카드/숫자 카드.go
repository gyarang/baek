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

	var N, M int
	fmt.Fscanln(reader, &N)
	cards := make(map[int]struct{})
	for i := 0; i < N; i++ {
		var card int
		fmt.Fscan(reader, &card)
		cards[card] = struct{}{}
	}
	fmt.Fscanln(reader)

	fmt.Fscanln(reader, &M)
	for i := 0; i < M; i++ {
		var input int
		fmt.Fscan(reader, &input)
		if _, ok := cards[input]; ok {
			fmt.Fprint(writer, "1 ")
		} else {
			fmt.Fprint(writer, "0 ")
		}
	}
}
