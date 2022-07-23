package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var l int
	fmt.Fscanf(rd, "%d\n", &l)

	var n int
	var str string
	var strArr []string

	for i := 0; i < l; i++ {
		fmt.Fscanf(rd, "%d %s\n", &n, &str)
		strArr = strings.Split(str, "")
		for _, s := range strArr {
			for j := 0; j < n; j++ {
				fmt.Fprint(wr, s)
			}
		}
		fmt.Fprint(wr, "\n")
	}
}
