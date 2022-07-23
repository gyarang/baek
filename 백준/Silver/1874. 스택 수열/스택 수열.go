package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var l, num, last int
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &l)

	stack := make([]int, 0, l)
	stack = append(stack, 0)
	for i := 0; i < l; i++ {
		fmt.Fscan(reader, &num)
		if num > stack[len(stack)-1] {
			for last < num {
				last += 1
				stack = append(stack, last)
				buf.WriteString("+\n")
			}
			stack = stack[:len(stack)-1]
			buf.WriteString("-\n")
		} else if num == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			buf.WriteString("-\n")
		} else {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print(buf.String())
}
