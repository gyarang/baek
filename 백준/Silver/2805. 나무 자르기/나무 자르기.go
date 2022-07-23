package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cnt, require int
	fmt.Fscanln(reader, &cnt, &require)

	arr := make([]int, cnt)
	max := 0
	for i := 0; i < cnt; i++ {
		fmt.Fscan(reader, &arr[i])
		if arr[i] > max {
			max = arr[i]
		}
	}

	start := 0
	end := max
	result := 0
	for start <= end {
		middle := (start + end) / 2
		sum := 0
		for _, v := range arr {
			if max-v > middle {
				continue
			} else {
				sum += middle - (max - v)
			}
		}

		if sum >= require {
			result = middle
			end = middle - 1
			continue
		} else {
			start = middle + 1
			continue
		}
	}
	fmt.Println(max - result)
}
