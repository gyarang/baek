package main

import (
	"bufio"
	"fmt"
	"os"
)

func lowerToUpper(r rune) rune {
	if r > 'Z' {
		return r - 32
	}
	return r
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	cnts := make(map[rune]int)

	var str string
	fmt.Fscanln(rd, &str)

	runes := []rune(str)
	for _, r := range runes {
		ur := lowerToUpper(r)
		if _, ok := cnts[ur]; !ok {
			cnts[ur] = 0
		}
		cnts[ur] += 1
	}

	var result rune
	var max int
	for key, val := range cnts {
		if max == val {
			result = '?'
		}
		if val > max {
			result = key
			max = val
		}
	}

	fmt.Fprintln(wr, string(result))
}
