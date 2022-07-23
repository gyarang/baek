package main

import (
	"bufio"
	"fmt"
	"os"
)

type Print struct {
	importance int
	observe    bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var testCase int
	fmt.Fscanln(reader, &testCase)

	var length, observe, importance int
TEST:
	for t := 0; t < testCase; t++ {
		fmt.Fscanln(reader, &length, &observe)
		queue := make([]Print, length)
		for i := 0; i < length; i++ {
			fmt.Fscanf(reader, "%d ", &importance)
			queue[i] = Print{importance, observe == i}
		}

		cnt := 0
	PRINT:
		for {
			q := queue[0]
			for _, v := range queue[1:] {
				if q.importance < v.importance {
					queue = append(queue[1:], queue[0])
					continue PRINT
				}
			}
			cnt++
			if q.observe {
				fmt.Println(cnt)
				continue TEST
			} else {
				queue = queue[1:]
			}
		}
	}
}
