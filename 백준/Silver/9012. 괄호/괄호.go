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

	var tc, cnt int
	var input string

	fmt.Fscanln(reader, &tc)
TEST:
	for i := 0; i < tc; i++ {
		cnt = 0
		fmt.Fscanln(reader, &input)
		for _, v := range input {
			if v == '(' {
				cnt++
			} else {
				cnt--
			}
			if cnt < 0 {
				fmt.Fprintln(writer, "NO")
				continue TEST
			}
		}
		if cnt == 0 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
