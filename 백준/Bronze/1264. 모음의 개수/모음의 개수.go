package main

import (
	"bufio"
	"fmt"
	"os"
)

func countVowels(input []byte) int {
	cnt := 0
	for _, v := range input {
		if v == 'a' || v == 'A' ||
			v == 'e' || v == 'E' ||
			v == 'i' || v == 'I' ||
			v == 'o' || v == 'O' ||
			v == 'u' || v == 'U' {
			cnt++
		}
	}
	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		input, _ := reader.ReadBytes('\n')
		if input[0] == '#' {
			return
		}
		fmt.Fprintln(writer, countVowels(input))
	}
}
