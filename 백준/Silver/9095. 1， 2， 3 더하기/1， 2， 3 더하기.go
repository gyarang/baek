package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	
	var cnt, tc int
	fmt.Fscanln(reader, &cnt)
	
	dpArr := make([]int, 12)
	dpArr[1] = 1
	dpArr[2] = 2
	dpArr[3] = 4
	last := 3

	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &tc)
		if tc < last {
			fmt.Fprintln(writer, dpArr[tc])
		} else {
			for i := last+1; i <= tc; i++ {
				dpArr[i] = dpArr[i-1] + dpArr[i-2] + dpArr[i-3]
			}
			fmt.Fprintln(writer, dpArr[tc])
			last = tc
		}
	}
}