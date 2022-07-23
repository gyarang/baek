package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l int
	fmt.Fscan(reader, &l)

	nmap := make(map[int]struct{})
	var num int
	for i := 0; i < l; i++ {
		fmt.Fscan(reader, &num)
		nmap[num] = struct{}{}
	}

	fmt.Fscan(reader, &l)

	var buf bytes.Buffer
	for i := 0; i < l; i++ {
		fmt.Fscan(reader, &num)
		if _, ok := nmap[num]; ok {
			buf.WriteString("1\n")
		} else {
			buf.WriteString("0\n")
		}
	}
	fmt.Print(buf.String())
}
