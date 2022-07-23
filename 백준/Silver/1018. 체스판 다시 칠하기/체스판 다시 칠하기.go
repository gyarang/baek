package main

import (
	"bufio"
	"fmt"
	"os"
)

func cntRepaint(plate [][]byte, w, h int) int {
	cnt := 0
	for i := w; i < w+8; i++ {
		for j := h; j < h+8; j++ {
			if (i+j)%2 == 0 && plate[j][i] != 'W' {
				cnt++
			} else if (i+j)%2 != 0 && plate[j][i] != 'B' {
				cnt++
			}
		}
	}

	if cnt > 32 {
		cnt = 64 - cnt
	}

	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var w, h int
	fmt.Fscanln(reader, &h, &w)

	plate := make([][]byte, h)
	for i := 0; i < h; i++ {
		var line string
		fmt.Fscanln(reader, &line)
		plate[i] = []byte(line)
	}

	min := 64
	for i := 0; i <= w-8; i++ {
		for j := 0; j <= h-8; j++ {
			p := cntRepaint(plate, i, j)
			if min > p {
				min = p
			}
		}
	}

	fmt.Println(min)
}
