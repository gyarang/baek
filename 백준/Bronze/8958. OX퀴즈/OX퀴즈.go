package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var l int
	fmt.Fscan(rd, &l)

	var str string
	var strArr []uint8
	var point int

	for i := 0; i < l; i++ {
		point = 0

		fmt.Fscan(rd, &str)
		strArr = []byte(str)
		p := 1
		for _, v := range strArr {
			if v == 'O' {
				point += p
				p++
			} else {
				p = 1
			}
		}
		fmt.Fprintln(wr, point)
	}
}
