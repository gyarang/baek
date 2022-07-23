package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var min, max int
	fmt.Scanln(&min, &max)

	arr := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		arr[i] = true
	}

	for i := 2; i*i <= max; i++ {
		if arr[i] {
			for j := i + i; j <= max; j += i {
				arr[j] = false
			}
		}
	}

	for i := min; i <= max; i++ {
		if arr[i] {
			fmt.Fprintln(writer, i)
		}
	}
}
