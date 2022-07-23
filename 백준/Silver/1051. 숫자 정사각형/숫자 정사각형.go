package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var w, h int
	fmt.Fscanln(reader, &h, &w)

	arr := make([][]int, h)
	for i := 0; i < h; i++ {
		var input string
		line := make([]int, w)
		fmt.Fscanln(reader, &input)
		for j, v := range input {
			line[j], _ = strconv.Atoi(string(v))
		}
		arr[i] = line
	}

	max := 1

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := j; k < w; k++ {
				if arr[i][j] == arr[i][k] {
					width := k - j
					if i+width < h {
						if arr[i][j] == arr[i+width][j] && arr[i][j] == arr[i+width][k] {
							size := (width + 1) * (width + 1)
							if size > max {
								max = size
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(max)
}
