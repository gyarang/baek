package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cnt, limit int
	fmt.Fscanln(reader, &cnt, &limit)

	dpArr := make([]int, limit + 1)

	for i := 0; i < cnt; i++ {
		var w, v int
		fmt.Fscanln(reader, &w, &v)
		for i := limit; i > w - 1; i-- {
			dpArr[i] = max(dpArr[i], dpArr[i-w]+v)
		}
	}
	
	fmt.Println(dpArr[limit])
}
