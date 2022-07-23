package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, sum, max int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &cards[i])
	}

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum = cards[i] + cards[j] + cards[k]
				if sum > max && sum <= m {
					max = sum
				}
			}
		}
	}

	fmt.Println(max)
}
