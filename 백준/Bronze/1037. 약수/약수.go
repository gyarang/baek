package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var count int
	fmt.Fscanln(reader, &count)

	min := 1000000
	max := 0

	for i := 0; i < count; i++ {
		var num int
		fmt.Fscanf(reader, "%d", &num)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println(min * max)
}
