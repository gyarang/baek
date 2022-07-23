package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a float64
	var b float64
	_, _ = fmt.Fscanf(r, "%f %f\n", &a, &b)

	fmt.Println(a/b)
}
