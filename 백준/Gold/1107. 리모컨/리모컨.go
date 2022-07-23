package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mod(a, b int) int {
	c := a - b
	if c > 0 {
		return c
	} else {
		return -1 * c
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var target, cnt, lost int
	fmt.Fscanln(reader, &target)
	fmt.Fscanln(reader, &cnt)

	var btns [10]bool
	for i := range btns {
		btns[i] = true
	}
	for i:=0; i < cnt; i++ {
		fmt.Fscan(reader, &lost)
		btns[lost] = false
	}

	min := mod(target, 100)
	Channel: for i := 0; i < 1000000; i++ {
		cnt := 0
		for _, v := range strconv.Itoa(i) {
			cnt++
			if !btns[v-'0'] {
				continue Channel
			}
			
		}

		result := mod(target, i) + cnt
		if min > result {
			min = result
		}
	}

	fmt.Println(min)
}
