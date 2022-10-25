package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsGroupWord(s string) bool {
	groupMap := make(map[byte]struct{})
	var prev byte
	for _, b := range []byte(s) {
		if b == prev {
			continue
		}
		if _, ok := groupMap[b]; ok {
			return false
		}
		prev = b
		groupMap[b] = struct{}{}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, result int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		input := ""
		fmt.Fscanln(reader, &input)
		if IsGroupWord(input) {
			result++
		}
	}

	fmt.Fprintln(writer, result)
}
