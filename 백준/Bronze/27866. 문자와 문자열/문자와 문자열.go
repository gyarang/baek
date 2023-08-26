package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input string
	var index int

	fmt.Fscanln(reader, &input)
	fmt.Fscanln(reader, &index)

	fmt.Println(string(input[index-1]))
}
