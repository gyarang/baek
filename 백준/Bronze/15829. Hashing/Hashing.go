package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	r = 31
	M = 1234567891
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var length int
	var str string
	fmt.Fscanln(reader, &length)
	fmt.Fscanln(reader, &str)

	var sum uint64
	for i, b := range str {
		num := uint64(b - 'a' + 1)
		for j := 0; j < i; j++ {
			num *= r
			num %= M
		}
		sum += num % M
	}

	fmt.Println(sum % M)
}
