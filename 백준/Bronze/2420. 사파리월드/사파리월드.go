package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d\n", Abs(a-b))
}
