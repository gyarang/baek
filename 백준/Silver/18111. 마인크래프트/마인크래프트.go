package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, b, max int
	fmt.Fscanln(reader, &n, &m, &b)

	land := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, m)
		for j := 0; j < m; j++ {
			var h int
			fmt.Fscan(reader, &h)
			if h > max {
				max = h
			}
			row[j] = h
		}
		land[i] = row
	}

	resultTime := 2147483647
	resultHeight := 0

	for i := max; i >= 0; i-- {
		var time, item int
		for _, row := range land {
			for _, spot := range row {
				if spot > i {
					time += (spot - i) * 2
					item += spot - i
				} else if spot < i {
					time += i - spot
					item -= i - spot
				}
			}
		}
		if b+item >= 0 && time < resultTime {
			resultTime = time
			resultHeight = i
		}
	}

	fmt.Println(resultTime, resultHeight)
}
