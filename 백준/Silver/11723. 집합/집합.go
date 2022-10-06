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

	cnt := 0
	fmt.Fscanln(reader, &cnt)
	set := 0

	for i := 0; i < cnt; i++ {
		what, val := "", 0
		fmt.Fscanln(reader, &what, &val)

		switch what {
		case "add":
			set = set | (1<<val)
		case "remove":
			set = set & (^(1<<val))
		case "toggle":
			set = set ^ (1<<val)
		case "all":
			set = 1 << 21 - 1
		case "empty":
			set = 0
		case "check":
			if set & (1 << val) == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, 1)
			}
		}
	}
}