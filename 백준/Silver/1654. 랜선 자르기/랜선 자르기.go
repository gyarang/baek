package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cnt, require, max int
	fmt.Fscanln(reader, &cnt, &require)

	arr := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &arr[i])
	}
	sort.Ints(arr)

	start := 1
	end := arr[cnt-1]

	for start <= end {
		mid := (start + end) / 2
		lineCnt := 0
		for _, v := range arr {
			lineCnt += v / mid
		}

		if lineCnt < require {
			end = mid - 1
			continue
		} else if lineCnt >= require {
			max = mid
			start = mid + 1
			continue
		}
	}
	fmt.Println(max)
}
