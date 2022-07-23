package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rd = bufio.NewReader(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	str, _ := rd.ReadString('\n')
	words := strings.Split(str, " ")
	var cnt int

	for _, v := range words {
		if v != "\n" && v != "" {
			cnt++
		}
	}

	fmt.Fprint(wr, cnt)
}
